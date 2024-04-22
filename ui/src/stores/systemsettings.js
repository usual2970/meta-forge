import { defineStore } from "pinia";
import {computed, ref} from "vue";
import {batchGet} from '@/api/systemsettings'


export const useSystemSettingsStore = defineStore("systemsettings", ()=>{
    const settings = ref({
        "@hasInitialized":0,
        schemas:{},
    });

    const hasInitialized=computed(()=>{
        return settings.value["@hasInitialized"];
    });
    const schemas=computed(()=>{
        return settings.value.schemas;
    });

    const getSettings=async ()=>{
        const resp=await batchGet({keys:"@hasInitialized,schemas"})
        if(resp.code!=0)
        {
            settings.value["@hasInitialized"]=-1;
        }else{
            if (!resp['data']['@hasInitialized']){
                settings.value["@hasInitialized"]=-1;
            }else{
                settings.value["@hasInitialized"]=1;
                
            }
            settings.value.schemas=!resp.data.schemas?{}:resp.data.schemas;
        }

        
        return settings.value;
    }


    return {settings,getSettings,hasInitialized,schemas}

});