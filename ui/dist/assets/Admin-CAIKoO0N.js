import{f as r,A as b,r as _,c as H,b as j,u as F,d as c,o as d,i as f,w as s,g as U,t as C,s as m,K as A,x as $}from"./index-G6iigneH.js";var B={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M408 442h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8zm-8 204c0 4.4 3.6 8 8 8h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56zm504-486H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zm0 632H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM115.4 518.9L271.7 642c5.8 4.6 14.4.5 14.4-6.9V388.9c0-7.4-8.5-11.5-14.4-6.9L115.4 505.1a8.74 8.74 0 000 13.8z"}}]},name:"menu-fold",theme:"outlined"};const D=B;function O(n){for(var t=1;t<arguments.length;t++){var e=arguments[t]!=null?Object(arguments[t]):{},a=Object.keys(e);typeof Object.getOwnPropertySymbols=="function"&&(a=a.concat(Object.getOwnPropertySymbols(e).filter(function(l){return Object.getOwnPropertyDescriptor(e,l).enumerable}))),a.forEach(function(l){L(n,l,e[l])})}return n}function L(n,t,e){return t in n?Object.defineProperty(n,t,{value:e,enumerable:!0,configurable:!0,writable:!0}):n[t]=e,n}var p=function(t,e){var a=O({},t,e.attrs);return r(b,O({},a,{icon:D}),null)};p.displayName="MenuFoldOutlined";p.inheritAttrs=!1;const N=p;var I={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M408 442h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8zm-8 204c0 4.4 3.6 8 8 8h480c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8H408c-4.4 0-8 3.6-8 8v56zm504-486H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zm0 632H120c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM142.4 642.1L298.7 519a8.84 8.84 0 000-13.9L142.4 381.9c-5.8-4.6-14.4-.5-14.4 6.9v246.3a8.9 8.9 0 0014.4 7z"}}]},name:"menu-unfold",theme:"outlined"};const V=I;function h(n){for(var t=1;t<arguments.length;t++){var e=arguments[t]!=null?Object(arguments[t]):{},a=Object.keys(e);typeof Object.getOwnPropertySymbols=="function"&&(a=a.concat(Object.getOwnPropertySymbols(e).filter(function(l){return Object.getOwnPropertyDescriptor(e,l).enumerable}))),a.forEach(function(l){E(n,l,e[l])})}return n}function E(n,t,e){return t in n?Object.defineProperty(n,t,{value:e,enumerable:!0,configurable:!0,writable:!0}):n[t]=e,n}var v=function(t,e){var a=h({},t,e.attrs);return r(b,h({},a,{icon:V}),null)};v.displayName="MenuUnfoldOutlined";v.inheritAttrs=!1;const K=v,R={class:"p-4 text-center text-white font-extrabold text-xl"},G={__name:"Admin",setup(n){const t=_(["1"]),e=_(!1),a=H(()=>e.value?"MF":"MetaForge"),M=j().menuItems,g=F(),x=u=>{if(u.keyPath.length==1&&g.push({name:u.key}),u.keyPath.length==2&&u.keyPath[0]=="entity"){g.push("/"+u.keyPath[0]+"/"+u.keyPath[1]);return}};return(u,o)=>{const P=c("a-menu"),k=c("a-layout-sider"),S=c("a-layout-header"),w=c("router-view"),z=c("a-layout-content"),y=c("a-layout");return d(),f(y,{id:"mf-layout",class:"font-mono"},{default:s(()=>[r(k,{collapsed:e.value,"onUpdate:collapsed":o[1]||(o[1]=i=>e.value=i),trigger:null,collapsible:""},{default:s(()=>[U("div",R,C(a.value),1),r(P,{selectedKeys:t.value,"onUpdate:selectedKeys":o[0]||(o[0]=i=>t.value=i),theme:"dark",mode:"inline",items:m(M),onClick:x},null,8,["selectedKeys","items"])]),_:1},8,["collapsed"]),r(y,null,{default:s(()=>[r(S,{style:{background:"#fff",padding:"0"}},{default:s(()=>[e.value?(d(),f(m(K),{key:0,class:"trigger",onClick:o[2]||(o[2]=()=>e.value=!e.value)})):(d(),f(m(N),{key:1,class:"trigger",onClick:o[3]||(o[3]=()=>e.value=!e.value)}))]),_:1}),r(z,{style:{margin:"24px 16px",padding:"24px",background:"#fff",minHeight:"280px"}},{default:s(()=>[r(w,null,{default:s(({Component:i})=>[(d(),f(A,null,[(d(),f($(i)))],1024))]),_:1})]),_:1})]),_:1})]),_:1})}}};export{G as default};