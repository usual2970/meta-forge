import{u as b,c as r,q as g,d as i,o as h,e as k,g as t,t as s,f as a,w as n,F as R,h as u,v as x}from"./index-G6iigneH.js";import{S}from"./SettingOutlined-C07MBp2x.js";const $={class:"flex justify-between"},O={class:"text-2xl text-slate-700"},w={class:"flex justify-between mt-2"},N={class:"mt-5 flex"},V=t("a",{href:"#"},"More",-1),B=t("p",null,"card content",-1),D={__name:"EntitySettingView",setup(j){const e=b(),c={crud:"增删改查",dict:"字典",field:"字段"},l=r(()=>c[e.currentRoute.value.params.type]+"设置"),d=r(()=>x(e.currentRoute.value.params.name)),m=r(()=>`/entity/${e.currentRoute.value.params.name}`),p=[{label:"增删改查",key:"basic",icon:()=>g(S),path:`/entity/${e.currentRoute.value.params.name}/setting/basic`},{type:"divider"},{label:"字典",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"字段",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"关联",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"关联",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"关联",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"关联",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`},{type:"divider"},{label:"关联",key:"field",icon:"SettingOutlined",path:`/entity/${e.currentRoute.value.params.name}/setting/field`}];return(C,E)=>{const _=i("router-link"),o=i("a-breadcrumb-item"),f=i("a-breadcrumb"),v=i("a-menu"),y=i("a-card");return h(),k(R,null,[t("div",$,[t("div",O,s(l.value),1)]),t("div",w,[a(f,{class:""},{default:n(()=>[a(o,null,{default:n(()=>[a(_,{to:m.value},{default:n(()=>[u(s(d.value),1)]),_:1},8,["to"])]),_:1}),a(o,null,{default:n(()=>[u(s(l.value),1)]),_:1})]),_:1})]),t("div",N,[t("div",null,[a(v,{items:p,class:"min-w-56 border rounded"})]),a(y,{title:l.value,bordered:!1,class:"ml-5 grow"},{extra:n(()=>[V]),default:n(()=>[B]),_:1},8,["title"])])],64)}}};export{D as default};
