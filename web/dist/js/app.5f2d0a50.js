(function(t){function e(e){for(var a,i,l=e[0],o=e[1],c=e[2],p=0,f=[];p<l.length;p++)i=l[p],Object.prototype.hasOwnProperty.call(r,i)&&r[i]&&f.push(r[i][0]),r[i]=0;for(a in o)Object.prototype.hasOwnProperty.call(o,a)&&(t[a]=o[a]);u&&u(e);while(f.length)f.shift()();return n.push.apply(n,c||[]),s()}function s(){for(var t,e=0;e<n.length;e++){for(var s=n[e],a=!0,l=1;l<s.length;l++){var o=s[l];0!==r[o]&&(a=!1)}a&&(n.splice(e--,1),t=i(i.s=s[0]))}return t}var a={},r={app:0},n=[];function i(e){if(a[e])return a[e].exports;var s=a[e]={i:e,l:!1,exports:{}};return t[e].call(s.exports,s,s.exports,i),s.l=!0,s.exports}i.m=t,i.c=a,i.d=function(t,e,s){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:s})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var s=Object.create(null);if(i.r(s),Object.defineProperty(s,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var a in t)i.d(s,a,function(e){return t[e]}.bind(null,a));return s},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/";var l=window["webpackJsonp"]=window["webpackJsonp"]||[],o=l.push.bind(l);l.push=e,l=l.slice();for(var c=0;c<l.length;c++)e(l[c]);var u=o;n.push([0,"chunk-vendors"]),s()})({0:function(t,e,s){t.exports=s("56d7")},"56d7":function(t,e,s){"use strict";s.r(e);s("cadf"),s("551c"),s("f751"),s("097d");var a=s("2b0e"),r=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"w-100 bg-light-red white avenir",attrs:{id:"app"}},[s("section",{staticClass:"mw5 mw7-ns center"},[s("nav",{staticClass:"db dt-l w-100 border-box pv4"}),t._m(0),s("section",{staticClass:"mw7 center"},[s("PopularArtists",{attrs:{msg:"Paste Your Taste for Last.FM"}})],1),t._m(1),t._m(2)])])},n=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("section",{staticClass:"mw7 center"},[s("h1",{staticClass:"f2"},[t._v("Paste My Taste 🎶")]),s("article",{staticClass:"lh-copy"},[s("p",[t._v("Paste My Taste generates an easily shareable list of top artists and genre information from your Last.FM account for a given timespan.")])])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("section",{staticClass:"mw7 center lh-copy"},[s("h2",{staticClass:"f3",attrs:{id:"about"}},[t._v("About 👋")]),t._v('\n            I built this replacement as the old "Paste My Taste" site went offline, if you have any feature requests or bug reports feel free to ping me on '),s("a",{staticClass:"link dim white underline",attrs:{href:"https://twitter.com/tehwey"}},[t._v("Twitter")]),t._v(".\n            "),s("p",[t._v("\n                If you want to support further development tell your friends or\n                "),s("a",{staticClass:"link dim white underline",attrs:{target:"_blank",href:"https://www.buymeacoffee.com/dewey"}},[s("span",{staticStyle:{"margin-right":"5px"}},[t._v("buy me a tea")]),s("img",{attrs:{src:"https://www.buymeacoffee.com/assets/img/BMC-btn-logo.svg",alt:"Buy me a tea"}})])])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("footer",{staticClass:"pv6 center white"},[s("p",{staticClass:"f6 db lh-solid"},[s("a",{staticClass:"pr2 link dim white",attrs:{href:"https://github.com/dewey"}},[t._v("Source on Github")]),s("a",{staticClass:"pr2 link dim white",attrs:{href:"https://twitter.com/tehwey"}},[t._v("Twitter")])])])}],i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("form",{staticClass:"mw7 center pa4 br2-ns",on:{submit:function(e){return e.preventDefault(),t.getPopularArtists.apply(null,arguments)}}},[s("fieldset",{staticClass:"cf bn ma0 pa0"},[s("div",{staticClass:"cf"},[s("div",{staticClass:"flex"},[s("div",{staticClass:"w-100 pa3"},[s("label",{staticClass:"clip",attrs:{for:"email-address"}},[t._v("Last.FM username")]),s("input",{directives:[{name:"model",rawName:"v-model",value:t.username,expression:"username"}],staticClass:"f6 f5-l tc input-reset bn fl black-80 bg-white pa3 w-100 br2-ns",attrs:{placeholder:"Last.FM Username",type:"text",name:"username",autocomplete:"off",autocorrect:"off",autocapitalize:"off",spellcheck:"false"},domProps:{value:t.username},on:{input:function(e){e.target.composing||(t.username=e.target.value)}}})])]),s("div",{staticClass:"flex pa3"},[s("div",{staticClass:"w-25 pr3"},[s("v-select",{attrs:{options:[{label:"Overall",value:"overall"},{label:"7 days",value:"7day"},{label:"2 weeks",value:"2week"},{label:"1 Month",value:"1month"},{label:"3 Months",value:"3month"},{label:"6 Months",value:"6month"},{label:"12 Months",value:"12month"}]},model:{value:t.period,callback:function(e){t.period=e},expression:"period"}})],1),s("div",{staticClass:"w-25 pr3"},[s("v-select",{attrs:{options:[{label:"10 artists",value:"10"},{label:"15 artists",value:"15"},{label:"20 artists",value:"20"},{label:"25 artists",value:"25"}]},model:{value:t.limit,callback:function(e){t.limit=e},expression:"limit"}})],1),s("div",{staticClass:"w-30 pr3"},[s("div",{staticClass:"flex items-center pv2"},[s("input",{directives:[{name:"model",rawName:"v-model",value:t.linkArtists,expression:"linkArtists"}],staticClass:"mr2 pa3",attrs:{type:"checkbox",id:"linkArtists",value:"false"},domProps:{checked:Array.isArray(t.linkArtists)?t._i(t.linkArtists,"false")>-1:t.linkArtists},on:{change:function(e){var s=t.linkArtists,a=e.target,r=!!a.checked;if(Array.isArray(s)){var n="false",i=t._i(s,n);a.checked?i<0&&(t.linkArtists=s.concat([n])):i>-1&&(t.linkArtists=s.slice(0,i).concat(s.slice(i+1)))}else t.linkArtists=r}}}),s("label",{staticClass:"lh-copy",attrs:{for:"linkArtists"}},[t._v("Add artist links")])])]),t._m(0)]),t.error.length>0?s("div",{staticClass:"flex"},[s("div",{staticClass:"w-100 pa3"},[t._v("\n            There was an error processing your request: "+t._s(t.error)+" 😔\n          ")])]):t._e()])])]),t.popularArtists.length>0?s("div",[s("h2",{staticClass:"f3",attrs:{id:"result"}},[t._v("Result 🥳")]),s("div",{staticClass:"pmt-result",on:{focus:function(t){return t.target.select()}}},[s("p",[t._v("I'm into "+t._s(t.popularTags)+" including:\n"+t._s(t.popularArtists)+"\n        ")]),s("p",[t._v("Check out my music taste: "),s("a",{attrs:{href:"https://www.last.fm/user/"+t.username}},[t._v("https://www.last.fm/user/"+t._s(t.username))])])])]):t._e()])},l=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"w-20"},[s("input",{staticClass:"f6 f5-l button-reset fl pv2 tc bn bg-animate bg-black-70 hover-bg-black white pointer w-100 br2-ns",attrs:{type:"submit",value:"Generate"}})])}],o=(s("7f7f"),s("ac6a"),s("bc3a")),c=s.n(o),u=function(){return c.a.create({baseURL:"/",headers:{Accept:"application/json","Content-Type":"application/json"}})},p={getPopularArtists:function(t){return u().get("/api/lastfm/"+t.username,{params:{limit:t.limit,period:t.period,linkArtists:t.linkArtists}}).then((function(t){return t.data}))}},f=s("6f79"),m=s.n(f),d=s("4a7a"),v=s.n(d),h={components:{vueSlider:m.a,vSelect:v.a},name:"PopularArtists",data:function(){return{username:"",period:{label:"Period"},limit:{label:"Limit"},linkArtists:!1,loading:!0,popularArtists:"",popularTags:"",error:""}},methods:{getPopularArtists:function(){var t=this;p.getPopularArtists({username:this.username,period:this.period.value,limit:this.limit.value}).then((function(e){var s=[],a=[],r="";e.forEach((function(e){t.linkArtists?s.push("[".concat(e.name,"](").concat(e.url,")")):s.push(e.name),a.push(e.genre)})),1===s.length?r=s[0]:2===s.length?r=s.join(" and "):s.length>2&&(r=s.slice(0,-1).join(", ")+" and "+s.slice(-1)+"."),t.popularArtists=r,t.popularTags=a.filter((function(t){return t})).join(", ")})).catch((function(e){return t.error=e.response.data})).finally((function(){t.loading=!1}))}}},b=h,w=(s("beff"),s("2877")),g=Object(w["a"])(b,i,l,!1,null,"b394de9e",null),y=g.exports,_={name:"app",components:{PopularArtists:y}},C=_,k=Object(w["a"])(C,r,n,!1,null,null,null),A=k.exports;s("6602");a["a"].config.productionTip=!1,new a["a"]({render:function(t){return t(A)}}).$mount("#app")},"729b":function(t,e,s){},beff:function(t,e,s){"use strict";s("729b")}});
//# sourceMappingURL=app.5f2d0a50.js.map