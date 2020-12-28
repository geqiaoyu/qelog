package manager

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/huzhongqing/qelog/pkg/httputil"

	"github.com/huzhongqing/qelog/pkg/common/entity"
	"github.com/huzhongqing/qelog/pkg/common/model"
	"github.com/huzhongqing/qelog/pkg/storage"
)

type Service struct {
	store *storage.Store
}

func NewService(store *storage.Store) *Service {
	srv := &Service{store: store}
	return srv
}

func (srv *Service) FindModuleList(ctx context.Context, in *entity.FindModuleListReq, out *entity.ListResp) error {
	filter := bson.M{}
	if in.Name != "" {
		filter["name"] = primitive.Regex{
			Pattern: in.Name,
			Options: "i",
		}
	}
	opt := options.Find()
	in.SetPage(opt)
	opt.SetSort(bson.M{"_id": -1})
	docs := make([]*model.Module, 0, in.PageSize)
	c, err := srv.store.FindModuleList(ctx, filter, &docs, opt)
	if err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}

	out.Count = c
	list := make([]*entity.FindModuleList, 0, len(docs))
	for _, v := range docs {
		d := &entity.FindModuleList{
			ID:             v.ID.Hex(),
			Name:           v.Name,
			Desc:           v.Desc,
			DBIndex:        v.DBIndex,
			HistoryDBIndex: v.HistoryDBIndex,
			UpdatedTsSec:   v.UpdatedAt.Unix(),
		}
		list = append(list, d)
	}
	out.List = list
	return nil
}

func (srv *Service) CreateModule(ctx context.Context, in *entity.CreateModuleReq) error {

	doc := &model.Module{
		Name:           in.Name,
		Desc:           in.Desc,
		DBIndex:        in.DBIndex,
		HistoryDBIndex: make([]int32, 0),
		UpdatedAt:      time.Now().Local(),
	}
	if err := srv.store.InsertModule(ctx, doc); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}
	return nil
}

func (srv *Service) UpdateModule(ctx context.Context, in *entity.UpdateModuleReq) error {
	id, err := in.ObjectID()
	if err != nil {
		return httputil.ErrArgsInvalid.MergeError(err)
	}

	doc := &model.Module{}
	if ok, err := srv.store.FindOneModule(ctx, bson.M{"_id": id}, doc); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	} else if !ok {
		return httputil.ErrNotFound
	}
	update := bson.M{}
	fields := bson.M{}
	if doc.DBIndex != in.DBIndex {
		fields["db_index"] = in.DBIndex
		update["$addToSet"] = bson.M{"history_db_index": in.DBIndex}
	}
	if doc.Desc != in.Desc {
		fields["desc"] = in.Desc
	}
	if len(fields) > 0 {
		fields["updated_at"] = time.Now().Local()
		update["$set"] = fields
	}
	if len(update) == 0 {
		return nil
	}
	filter := bson.M{
		"_id":        doc.ID,
		"updated_at": doc.UpdatedAt,
	}
	return srv.store.UpdateModule(ctx, filter, update)
}

func (srv *Service) DeleteModule(ctx context.Context, in *entity.DeleteModuleReq) error {
	id, err := in.ObjectID()
	if err != nil {
		return httputil.ErrArgsInvalid.MergeError(err)
	}
	doc := &model.Module{}
	if ok, err := srv.store.FindOneModule(ctx, bson.M{"_id": id}, doc); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	} else if !ok {
		return httputil.ErrNotFound
	}
	if doc.Name != in.Name {
		return httputil.ErrNotFound
	}
	return srv.store.DeleteModule(ctx, id)
}

func (srv *Service) FindLoggingList(ctx context.Context, in *entity.FindLoggingListReq, out *entity.ListResp) error {

	// 如果没有传入时间，则默认设置一个间隔时间
	b, e := in.InitTimeSection(time.Hour)
	// 计算查询时间应该在哪个分片
	collectionName := model.LoggingCollectionName(in.DBIndex, b.Unix())
	if collectionName != model.LoggingCollectionName(in.DBIndex, e.Unix()) {
		return httputil.ErrArgsInvalid.MergeError(fmt.Errorf("查询时间跨度不能超过时间分片设置 (分片粒度 %s)", model.LoggingShardingTime))
	}

	filter := bson.M{
		"m": in.ModuleName,
	}

	// 必须存在时间
	filter["ts"] = bson.M{"$gte": b.Unix(), "$lt": e.Unix()}

	if in.Short != "" {
		// 区分大小写
		filter["s"] = primitive.Regex{
			Pattern: in.Short,
		}
	}

	if in.Level >= 0 {
		filter["l"] = in.Level
	}
	// 必需要有前置条件，才能查询后面的，以便命中索引
	if in.ConditionOne != "" {
		filter["c1"] = in.ConditionOne
		if in.ConditionTwo != "" {
			filter["c2"] = in.ConditionTwo
			if in.ConditionThree != "" {
				filter["c3"] = in.ConditionThree
			}
		}
	}

	findOpt := options.Find()
	in.SetPage(findOpt)
	findOpt.SetSort(bson.M{"ts": -1})

	docs := make([]*model.Logging, 0, in.PageSize)
	c, err := srv.store.FindLoggingList(ctx, collectionName, filter, &docs, findOpt)
	if err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}

	out.Count = c

	// 去除极低可能重复写入的日志信息
	hitMap := map[string]struct{}{}
	list := make([]*entity.FindLoggingList, 0, len(docs))
	for _, v := range docs {
		if _, ok := hitMap[v.MessageID]; ok {
			continue
		} else {
			hitMap[v.MessageID] = struct{}{}
		}

		d := &entity.FindLoggingList{
			ID:             v.ID.Hex(),
			TsMill:         v.Time,
			Level:          int32(v.Level),
			ShortMsg:       v.Short,
			Full:           v.Full,
			ConditionOne:   v.Condition1,
			ConditionTwo:   v.Condition2,
			ConditionThree: v.Condition3,
			IP:             v.IP,
		}
		list = append(list, d)
	}
	out.List = list

	return nil
}

type AscDBIndexState []entity.DBIndexState

func (v AscDBIndexState) Len() int           { return len(v) }
func (v AscDBIndexState) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v AscDBIndexState) Less(i, j int) bool { return v[i].Use < v[j].Use }

func (srv *Service) GetDBIndex(ctx context.Context, out *entity.GetDBIndexResp) error {
	docs, err := srv.store.FindAllModule(ctx)
	if err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}
	state := make(map[int32]int)
	for i := int32(1); i <= model.MaxDBIndex; i++ {
		state[i] = 0
	}
	for _, v := range docs {
		num, ok := state[v.DBIndex]
		if ok {
			state[v.DBIndex] = num + 1
		}
	}

	states := make([]entity.DBIndexState, 0, len(state))
	for k, v := range state {
		states = append(states, entity.DBIndexState{
			Index: k,
			Use:   v,
		})
	}
	sort.Sort(AscDBIndexState(states))

	// 找到最小的，作为推荐
	suggestDBIndex := model.MaxDBIndex
	if len(states) > 0 {
		suggestDBIndex = states[0].Index
	}

	out.SuggestDBIndex = suggestDBIndex
	out.MaxDBIndex = model.MaxDBIndex
	out.UseState = states

	return nil
}

func (srv *Service) FindAlarmRuleList(ctx context.Context, in *entity.FindAlarmRuleListReq, out *entity.ListResp) error {
	filter := bson.M{}
	if in.ModuleName != "" {
		filter["module_name"] = primitive.Regex{
			Pattern: in.ModuleName,
			Options: "i",
		}
	}
	if in.Enable > 0 {
		filter["enable"] = in.Enable == 1
	}
	if in.Short != "" {
		filter["short"] = in.Short
	}

	opt := options.Find()
	in.SetPage(opt)
	opt.SetSort(bson.M{"_id": -1})
	docs := make([]*model.AlarmRule, 0, in.PageSize)
	c, err := srv.store.FindAlarmRuleList(ctx, filter, &docs, opt)
	if err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}

	out.Count = c
	list := make([]*entity.FindAlarmRuleList, 0, len(docs))
	for _, v := range docs {
		d := &entity.FindAlarmRuleList{
			ID:           v.ID.Hex(),
			Enable:       v.Enable,
			ModuleName:   v.ModuleName,
			Short:        v.Short,
			Level:        v.Level.Int32(),
			Tag:          v.Tag,
			RateSec:      v.RateSec,
			Method:       v.Method,
			HookURL:      v.HookURL,
			UpdatedTsSec: v.UpdatedAt.Unix(),
		}
		list = append(list, d)
	}
	out.List = list

	return nil
}

func (srv *Service) CreateAlarmRule(ctx context.Context, in *entity.CreateAlarmRuleReq) error {

	doc := &model.AlarmRule{
		Enable:     true,
		ModuleName: in.ModuleName,
		Short:      in.Short,
		Level:      model.Level(in.Level),
		Tag:        in.Tag,
		RateSec:    in.RateSec,
		Method:     in.Method,
		HookURL:    in.HookURL,
		UpdatedAt:  time.Now().Local(),
	}

	if err := srv.store.InsertAlarmRule(ctx, doc); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}

	return nil
}

func (srv *Service) UpdateAlarmRule(ctx context.Context, in *entity.UpdateAlarmRuleReq) error {
	id, err := in.ObjectID()
	if err != nil {
		return httputil.ErrArgsInvalid.MergeError(err)
	}

	doc := &model.AlarmRule{}
	if ok, err := srv.store.FindOneAlarmRule(ctx, bson.M{"_id": id}, doc); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	} else if !ok {
		return httputil.ErrNotFound
	}
	update := bson.M{}
	fields := bson.M{}
	if (in.Enable == 1) != doc.Enable {
		fields["enable"] = in.Enable == 1
	}
	if in.Short != doc.Short {
		fields["short"] = in.Short
	}
	if in.RateSec != doc.RateSec {
		fields["rate_sec"] = in.RateSec
	}
	if in.Level != doc.Level.Int32() {
		fields["level"] = in.Level
	}
	if in.Tag != doc.Tag {
		fields["tag"] = in.Tag
	}
	if in.Method != doc.Method {
		fields["method"] = in.Method
	}
	if in.HookURL != doc.HookURL {
		fields["hook_url"] = in.HookURL
	}

	if len(fields) > 0 {
		fields["updated_at"] = time.Now().Local()
		update["$set"] = fields
	}
	if len(update) == 0 {
		return nil
	}

	filter := bson.M{
		"_id":        id,
		"updated_at": doc.UpdatedAt,
	}

	if err := srv.store.UpdateAlarmRule(ctx, filter, update); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}

	return nil
}

func (srv *Service) DeleteAlarmRule(ctx context.Context, in *entity.DeleteAlarmRuleReq) error {
	id, err := in.ObjectID()
	if err != nil {
		return httputil.ErrArgsInvalid.MergeError(err)
	}
	if err := srv.store.DeleteAlarmRule(ctx, id); err != nil {
		return httputil.ErrSystemException.MergeError(err)
	}
	return nil
}
