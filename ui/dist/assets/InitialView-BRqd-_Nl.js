import{r as N,a as V,c,u as B,b as F,d as n,o as i,e as _,f as t,w as s,g as d,h as m,i as k,j as f,F as z,k as R,m as h}from"./index-BsyXZ6rX.js";const I={class:"min-w-[500px] border rounded-md bg-blue-500 font-mono"},M=d("div",{class:"p-7"},[d("div",{class:"h-11 text-4xl text-center text-white"},"MetaForge"),d("div",{class:"text-xl pt-1 text-center text-white"},"数据库配置"),d("div",{class:"text-sm text-center text-white pt-1"},"设置数据库连接")],-1),O={__name:"InitialView",setup(E){const v=N(),a=V({kind:"",file:"",host:"",port:"",user:"",password:"",database:""}),b=c(()=>a.kind==="sqlite"),g=c(()=>["mysql","postgresql"].includes(a.kind)),y=c(()=>b.value||g.value),w={kind:[{message:"Please select a kind",required:!0,trigger:"change"}],file:[{validator:async(r,e)=>a.kind=="sqlite"&&e==""?Promise.reject("Please input a file"):Promise.resolve(),trigger:"change"}],host:[{validator:async(r,e)=>a.kind=="mysql"&&e==""?Promise.reject("Please input a host"):Promise.resolve(),trigger:"change"}],port:[{validator:async(r,e)=>a.kind=="mysql"&&e==""?Promise.reject("Please input a port"):Promise.resolve(),trigger:"change"}],user:[{validator:async(r,e)=>a.kind=="mysql"&&e==""?Promise.reject("Please input a user"):Promise.resolve(),trigger:"change"}],password:[{validator:async(r,e)=>a.kind=="mysql"&&e==""?Promise.reject("Please input a password"):Promise.resolve(),trigger:"change"}],database:[{validator:async(r,e)=>a.kind=="mysql"&&e==""?Promise.reject("Please input a database"):Promise.resolve(),trigger:"change"}]},P=B(),x=F(),q=async()=>{try{await v.value.validate();const r=await R(a);return r.code==0?(h.success("Initialize success",2),await x.getSettings(),P.push({name:"home"})):h.error(r.msg)}catch(r){console.log(JSON.stringify(r))}};return(r,e)=>{const p=n("a-select-option"),C=n("a-select"),o=n("a-form-item"),u=n("a-input"),S=n("a-input-password"),U=n("a-button"),j=n("a-form");return i(),_("div",I,[M,t(j,{class:"p-7 bg-white",labelCol:{span:5},ref_key:"formRef",ref:v,rules:w,model:a},{default:s(()=>[t(o,{label:"数据库类型","has-feedback":"",name:"kind"},{default:s(()=>[t(C,{value:a.kind,"onUpdate:value":e[0]||(e[0]=l=>a.kind=l),placeholder:"数据库类型"},{default:s(()=>[t(p,{value:""},{default:s(()=>[m("--选择数据库类型--")]),_:1}),t(p,{value:"mysql"},{default:s(()=>[m("Mysql")]),_:1}),t(p,{value:"sqlite"},{default:s(()=>[m("Sqlite")]),_:1})]),_:1},8,["value"])]),_:1}),b.value?(i(),k(o,{key:0,label:"数据库文件",labelCol:{span:5},"has-feedback":"",name:"file"},{default:s(()=>[t(u,{value:a.file,"onUpdate:value":e[1]||(e[1]=l=>a.file=l)},null,8,["value"])]),_:1})):f("",!0),g.value?(i(),_(z,{key:1},[t(o,{label:"主机(host)",labelCol:{span:5},"has-feedback":"",name:"host"},{default:s(()=>[t(u,{value:a.host,"onUpdate:value":e[2]||(e[2]=l=>a.host=l)},null,8,["value"])]),_:1}),t(o,{label:"端口(port)",labelCol:{span:5},"has-feedback":"",name:"port"},{default:s(()=>[t(u,{value:a.port,"onUpdate:value":e[3]||(e[3]=l=>a.port=l)},null,8,["value"])]),_:1}),t(o,{label:"用户",labelCol:{span:5},"has-feedback":"",name:"user"},{default:s(()=>[t(u,{value:a.user,"onUpdate:value":e[4]||(e[4]=l=>a.user=l)},null,8,["value"])]),_:1}),t(o,{label:"密码",labelCol:{span:5},"has-feedback":"",name:"password"},{default:s(()=>[t(S,{value:a.password,"onUpdate:value":e[5]||(e[5]=l=>a.password=l)},null,8,["value"])]),_:1}),t(o,{label:"数据库",labelCol:{span:5},"has-feedback":"",name:"database"},{default:s(()=>[t(u,{value:a.database,"onUpdate:value":e[6]||(e[6]=l=>a.database=l)},null,8,["value"])]),_:1})],64)):f("",!0),y.value?(i(),k(o,{key:2,"wrapper-col":{offset:19}},{default:s(()=>[t(U,{type:"primary","html-type":"submit",size:"large",class:"bg-blue-500",onClick:q},{default:s(()=>[m("提交")]),_:1})]),_:1})):f("",!0)]),_:1},8,["model"])])}}};export{O as default};
