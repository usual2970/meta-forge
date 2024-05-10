
import {useSystemSettingsStore} from '@/stores/systemsettings'
export async function initializeMiddleware(to,_from) {
    
    // 获取初始化状态
    const store=useSystemSettingsStore()
    if (store.hasInitialized==0){
        await store.getSettings()
    }
    
    // 如果未初始化，且不是initial页，跳转到initial页
    if (store.hasInitialized==-1 && to.name!="initial"){
        return {name:"initial"}
    }

    // 如果已初始化，且是initial页，跳转到home页
    if (store.hasInitialized==1 && to.name=='initial'){
        return {name:"home"}
    }

    // 第一个mvp不要首页
    if (to.name=='home'){
        return {name:"entity", params:{name:store.schemas[0].name}}
    }

    return true

}