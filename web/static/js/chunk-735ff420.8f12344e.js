(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-735ff420"],{"333d":function(e,t,a){"use strict";var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"pagination-container",class:{hidden:e.hidden}},[a("el-pagination",e._b({attrs:{background:e.background,"current-page":e.currentPage,"page-size":e.pageSize,layout:e.layout,"page-sizes":e.pageSizes,total:e.total},on:{"update:currentPage":function(t){e.currentPage=t},"update:current-page":function(t){e.currentPage=t},"update:pageSize":function(t){e.pageSize=t},"update:page-size":function(t){e.pageSize=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}},"el-pagination",e.$attrs,!1))],1)},n=[];a("c5f6");Math.easeInOutQuad=function(e,t,a,l){return e/=l/2,e<1?a/2*e*e+t:(e--,-a/2*(e*(e-2)-1)+t)};var i=function(){return window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||function(e){window.setTimeout(e,1e3/60)}}();function r(e){document.documentElement.scrollTop=e,document.body.parentNode.scrollTop=e,document.body.scrollTop=e}function o(){return document.documentElement.scrollTop||document.body.parentNode.scrollTop||document.body.scrollTop}function u(e,t,a){var l=o(),n=e-l,u=20,s=0;t="undefined"===typeof t?500:t;var c=function e(){s+=u;var o=Math.easeInOutQuad(s,l,n,t);r(o),s<t?i(e):a&&"function"===typeof a&&a()};c()}var s={name:"Pagination",props:{total:{required:!0,type:Number},page:{type:Number,default:1},limit:{type:Number,default:20},pageSizes:{type:Array,default:function(){return[10,20,30,50]}},layout:{type:String,default:"total, sizes, prev, pager, next, jumper"},background:{type:Boolean,default:!0},autoScroll:{type:Boolean,default:!0},hidden:{type:Boolean,default:!1}},computed:{currentPage:{get:function(){return this.page},set:function(e){this.$emit("update:page",e)}},pageSize:{get:function(){return this.limit},set:function(e){this.$emit("update:limit",e)}}},methods:{handleSizeChange:function(e){this.$emit("pagination",{page:this.currentPage,limit:e}),this.autoScroll&&u(0,800)},handleCurrentChange:function(e){this.$emit("pagination",{page:e,limit:this.pageSize}),this.autoScroll&&u(0,800)}}},c=s,d=(a("e498"),a("2877")),m=Object(d["a"])(c,l,n,!1,null,"6af373ef",null);t["a"]=m.exports},7456:function(e,t,a){},a372:function(e,t,a){"use strict";a.d(t,"q",(function(){return n})),a.d(t,"p",(function(){return i})),a.d(t,"c",(function(){return r})),a.d(t,"u",(function(){return o})),a.d(t,"g",(function(){return u})),a.d(t,"h",(function(){return s})),a.d(t,"a",(function(){return c})),a.d(t,"s",(function(){return d})),a.d(t,"d",(function(){return m})),a.d(t,"i",(function(){return f})),a.d(t,"b",(function(){return p})),a.d(t,"t",(function(){return h})),a.d(t,"e",(function(){return g})),a.d(t,"r",(function(){return v})),a.d(t,"k",(function(){return b})),a.d(t,"j",(function(){return y})),a.d(t,"f",(function(){return k})),a.d(t,"m",(function(){return _})),a.d(t,"l",(function(){return S})),a.d(t,"n",(function(){return x})),a.d(t,"o",(function(){return R}));var l=a("b775");function n(e){return Object(l["a"])({url:"/shardingIndex",method:"get"})}function i(e){return Object(l["a"])({url:"/module/list",method:"get",params:e})}function r(e){return Object(l["a"])({url:"/module",method:"post",data:e})}function o(e){return Object(l["a"])({url:"/module",method:"put",data:e})}function u(e){return Object(l["a"])({url:"/module",method:"delete",data:e})}function s(e){return Object(l["a"])({url:"/alarmRule/list",method:"get",params:e})}function c(e){return Object(l["a"])({url:"/alarmRule",method:"post",data:e})}function d(e){return Object(l["a"])({url:"/alarmRule",method:"put",data:e})}function m(e){return Object(l["a"])({url:"/alarmRule",method:"delete",data:e})}function f(e){return Object(l["a"])({url:"/alarmRule/hook/list",method:"get",params:e})}function p(e){return Object(l["a"])({url:"/alarmRule/hook",method:"post",data:e})}function h(e){return Object(l["a"])({url:"/alarmRule/hook",method:"put",data:e})}function g(e){return Object(l["a"])({url:"/alarmRule/hook",method:"delete",data:e})}function v(e){return Object(l["a"])({url:"/alarmRule/hook/ping",method:"get",params:e})}function b(e){return Object(l["a"])({url:"/logging/list",method:"post",data:e})}function y(e){return Object(l["a"])({url:"/logging/traceid",method:"post",data:e})}function k(e){return Object(l["a"])({url:"/logging/collection",method:"delete",data:e})}function _(){return Object(l["a"])({url:"/metrics/dbStats",method:"get"})}function S(e){return Object(l["a"])({url:"/metrics/collStats",method:"get",params:e})}function x(e){return Object(l["a"])({url:"/metrics/module/list",method:"get",params:e})}function R(e){return Object(l["a"])({url:"/metrics/module/trend",method:"get",params:e})}},cc38:function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("div",{staticClass:"filter-container"},[a("el-input",{staticClass:"filter-item",staticStyle:{width:"200px"},attrs:{placeholder:"模块名称"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleFilter(t)}},model:{value:e.listQuery.moduleName,callback:function(t){e.$set(e.listQuery,"moduleName",t)},expression:"listQuery.moduleName"}}),e._v(" "),a("el-input",{staticClass:"filter-item",staticStyle:{width:"200px"},attrs:{placeholder:"短消息"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleFilter(t)}},model:{value:e.listQuery.short,callback:function(t){e.$set(e.listQuery,"short",t)},expression:"listQuery.short"}}),e._v(" "),a("el-select",{staticClass:"filter-item",attrs:{placeholder:"状态"},model:{value:e.listQuery.enable,callback:function(t){e.$set(e.listQuery,"enable",t)},expression:"listQuery.enable"}},e._l(e.enableSorts,(function(e){return a("el-option",{key:e.index,attrs:{label:e.value,value:e.index}})})),1),e._v(" "),a("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.handleFilter}},[e._v("\n      搜索\n    ")]),e._v(" "),a("el-button",{staticClass:"filter-item",staticStyle:{"margin-left":"10px"},attrs:{type:"primary",icon:"el-icon-edit"},on:{click:e.handleCreate}},[e._v("\n      新增\n    ")])],1),e._v(" "),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],staticStyle:{width:"100%"},attrs:{data:e.list,border:"",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{label:"更新时间",width:"150px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s(e._f("parseTime")(l.updatedTsSec,"{y}-{m}-{d} {h}:{i}")))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"状态",prop:"enabel",align:"center",width:"100px"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("el-tag",{attrs:{type:e._f("enableFilter")(l.enable)}},[e._v("\n          "+e._s(l.enable)+"\n        ")])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"模块名",prop:"moduleName",align:"center",width:"150px"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s(l.moduleName))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"短消息","min-width":"100px"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s(l.short))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"等级",prop:"enabel",align:"center",width:"100px"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[e._v("\n        "+e._s(e.sortsFilter(a.level,"levelSorts"))+"\n      ")]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"标签","min-width":"50px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s(l.tag))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"频率","min-width":"50px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s("1/"+l.rateSec+"s"))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"通知方式","min-width":"50px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(" "+e._s(e.sortsFilter(l.method,"methodSorts")))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"HookURL","min-width":"50px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row;return[a("span",[e._v(e._s(e.hookNameFilter(l.hookId)))])]}}])}),e._v(" "),a("el-table-column",{attrs:{label:"操作",align:"center",width:"230","class-name":"small-padding fixed-width"},scopedSlots:e._u([{key:"default",fn:function(t){var l=t.row,n=t.$index;return[a("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(t){return e.handleUpdate(l)}}},[e._v("\n          编辑\n        ")]),e._v(" "),"deleted"!=l.status?a("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(t){return e.handleDelete(l,n)}}},[e._v("\n          删除\n        ")]):e._e()]}}])})],1),e._v(" "),a("pagination",{directives:[{name:"show",rawName:"v-show",value:e.total>0,expression:"total > 0"}],attrs:{total:e.total,page:e.listQuery.page,limit:e.listQuery.limit},on:{"update:page":function(t){return e.$set(e.listQuery,"page",t)},"update:limit":function(t){return e.$set(e.listQuery,"limit",t)},pagination:e.getList}}),e._v(" "),a("el-dialog",{attrs:{title:e.textMap[e.dialogStatus],visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[a("el-form",{ref:"dataForm",staticStyle:{width:"400px","margin-left":"50px"},attrs:{model:e.alarmRule,rules:"delete"!==e.dialogStatus?e.rules:{},disabled:"delete"===e.dialogStatus,"label-position":"left","label-width":"120px"}},[a("el-form-item",{attrs:{label:"模块名",prop:"moduleName"}},[a("el-select",{staticClass:"filter-item",staticStyle:{width:"100%"},attrs:{filterable:"",placeholder:"请选择模块名",disabled:"create"!=e.dialogStatus},model:{value:e.alarmRule.moduleName,callback:function(t){e.$set(e.alarmRule,"moduleName",t)},expression:"alarmRule.moduleName"}},e._l(e.modules,(function(e){return a("el-option",{key:e.id,attrs:{label:e.moduleName,value:e.moduleName}})})),1)],1),e._v(" "),a("el-form-item",{attrs:{label:"短消息",prop:"short"}},[a("el-input",{model:{value:e.alarmRule.short,callback:function(t){e.$set(e.alarmRule,"short",t)},expression:"alarmRule.short"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"等级",prop:"level","min-width":"80px"}},[a("el-select",{staticClass:"filter-item",attrs:{placeholder:"Please select"},model:{value:e.alarmRule.level,callback:function(t){e.$set(e.alarmRule,"level",t)},expression:"alarmRule.level"}},e._l(e.levelSorts,(function(e){return a("el-option",{key:e.index,attrs:{label:e.value,value:e.index}})})),1)],1),e._v(" "),a("el-form-item",{attrs:{label:"通知方式",prop:"method",width:"110px"}},[a("el-select",{staticClass:"filter-item",attrs:{placeholder:"Please select"},model:{value:e.alarmRule.method,callback:function(t){e.$set(e.alarmRule,"method",t)},expression:"alarmRule.method"}},e._l(e.methodSorts,(function(e){return a("el-option",{key:e.index,attrs:{label:e.value,value:e.index}})})),1)],1),e._v(" "),a("el-form-item",{attrs:{label:"HookURL"}},[a("el-select",{staticClass:"filter-item",staticStyle:{width:"100%"},attrs:{filterable:"",placeholder:"请选择HookURL"},model:{value:e.alarmRule.hookId,callback:function(t){e.$set(e.alarmRule,"hookId",t)},expression:"alarmRule.hookId"}},e._l(e.hooks,(function(e){return a("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1)],1),e._v(" "),a("el-form-item",{attrs:{label:"频率间隔(秒)",prop:"rateSec"}},[a("el-input-number",{model:{value:e.alarmRule.rateSec,callback:function(t){e.$set(e.alarmRule,"rateSec",t)},expression:"alarmRule.rateSec"}})],1),e._v(" "),"update"==e.dialogStatus?a("el-form-item",{attrs:{label:"关闭/启动",prop:"enable"}},[a("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949"},model:{value:e.alarmRule.enable,callback:function(t){e.$set(e.alarmRule,"enable",t)},expression:"alarmRule.enable"}})],1):e._e(),e._v(" "),a("el-form-item",{attrs:{label:"标签",prop:"tag"}},[a("el-input",{model:{value:e.alarmRule.tag,callback:function(t){e.$set(e.alarmRule,"tag",t)},expression:"alarmRule.tag"}})],1)],1),e._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v(" 取消 ")]),e._v(" "),a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleConfrim(e.dialogStatus)}}},[e._v("\n        确认\n      ")])],1)],1)],1)},n=[],i=(a("ac6a"),a("7f7f"),a("a372")),r=a("333d"),o={name:"Alarm",components:{Pagination:r["a"]},filters:{enableFilter:function(e){return e?"success":"danger"}},data:function(){return{list:null,total:0,listLoading:!0,listQuery:{page:1,limit:20,moduleName:"",enable:-1,short:""},enableSorts:[{index:-1,value:"状态"},{index:0,value:"关闭"},{index:1,value:"开启"}],levelSorts:[{index:-1,value:"DEBUG"},{index:0,value:"INFO"},{index:1,value:"WARN"},{index:2,value:"ERROR"},{index:3,value:"DPANIC"},{index:4,value:"PANIC"},{index:5,value:"FATAL"}],methodSorts:[{index:1,value:"DingDing"}],alarmRule:{id:void 0,moduleName:"",short:"",level:0,tag:"",rateSec:30,method:1,hookId:"",enabel:-1},dialogFormVisible:!1,dialogStatus:"",textMap:{update:"编辑",create:"创建",delete:"删除"},hooks:[],hooksMap:{},modules:[],rules:{moduleName:[{required:!0,message:"moduleName is required",trigger:"change"}],short:[{required:!0,message:"short is required",trigger:"change"}],hookId:[{required:!0,message:"hookId is required",trigger:"change"}]}}},created:function(){this.getList(),this.fetchHookURLList(),this.fetchModuleList()},methods:{sortsFilter:function(e,t){var a=this[t];if(a)for(var l=0;l<a.length;l++){var n=a[l],i=n.index,r=n.value;if(i===e)return r}},hookNameFilter:function(e){if(e&&this.hooksMap){var t=this.hooksMap[e];if(t)return t.name}return""},getList:function(){var e=this;this.listLoading=!0,Object(i["h"])(this.listQuery).then((function(t){e.list=t.data.list,e.total=t.data.count,setTimeout((function(){e.listLoading=!1}),500)}))},handleFilter:function(){this.listQuery.page=1,this.getList()},resetAlarmRule:function(){this.alarmRule={id:void 0,moduleName:"",short:"",level:-1,tag:"",rateSec:30,method:1,hookId:"",enabel:!0}},handleCreate:function(){var e=this;this.resetAlarmRule(),this.dialogStatus="create",this.dialogFormVisible=!0,this.$nextTick((function(){e.$refs["dataForm"].clearValidate()}))},fetchHookURLList:function(){var e=this;Object(i["i"])({method:this.alarmRule.method,page:1,limit:50}).then((function(t){e.listLoading=!1;var a=t.data.list,l=[],n={};a.forEach((function(e){var t={id:e.id,name:e.name,url:e.url,method:e.method};n[e.id]=t,l.push(t)})),e.hooks=l,e.hooksMap=n}))},fetchModuleList:function(){var e=this;Object(i["p"])({page:1,limit:150}).then((function(t){e.listLoading=!1;var a=t.data.list,l=[];a.forEach((function(e){var t={id:e.id,moduleName:e.name};l.push(t)})),e.modules=l}))},createData:function(){var e=this;this.$refs["dataForm"].validate((function(t){t&&Object(i["a"])(e.alarmRule).then((function(){e.getList(),e.dialogFormVisible=!1,e.$notify({title:"Success",message:"新增成功",type:"success",duration:2e3})}))}))},handleUpdate:function(e){var t=this;this.alarmRule=Object.assign({},e),console.log(this.alarmRule),this.dialogStatus="update",this.dialogFormVisible=!0,this.$nextTick((function(){t.$refs["dataForm"].clearValidate()}))},updateData:function(){var e=this;this.$refs["dataForm"].validate((function(t){t&&Object(i["s"])(e.alarmRule).then((function(){e.getList(),e.dialogFormVisible=!1,e.$notify({title:"Success",message:"编辑成功",type:"success",duration:2e3})}))}))},handleDelete:function(e,t){var a=this;this.alarmRule=Object.assign({},e),this.dialogStatus="delete",this.dialogFormVisible=!0,this.$nextTick((function(){a.$refs["dataForm"].clearValidate()}))},deleteData:function(){var e=this;this.$refs["dataForm"].validate((function(t){t&&Object(i["d"])({id:e.alarmRule.id}).then((function(){e.getList(),e.dialogFormVisible=!1,e.$notify({title:"Success",message:"删除成功",type:"success",duration:2e3})}))}))},handleConfrim:function(e){switch(e){case"create":return this.createData();case"update":return this.updateData();case"delete":return this.deleteData()}}}},u=o,s=a("2877"),c=Object(s["a"])(u,l,n,!1,null,null,null);t["default"]=c.exports},e498:function(e,t,a){"use strict";a("7456")}}]);