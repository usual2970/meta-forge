import { defineStore } from "pinia";
import {computed, ref} from "vue";


export const useSystemSettingsStore = defineStore("systemsettings", ()=>{
    const settings = ref({
        "@hasInitialized":0,
        schemas:{},
    });

    const hasInitialized=computed(()=>{
        return settings.value["@hasInitialized"]!=0;
    });
    const schemas=computed(()=>{
        return settings.value.schemas;
    });

    const getSettings=async ()=>{
        return settings.value;
    }

});