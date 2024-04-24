import { defineStore } from "pinia";
import {computed, ref,h} from "vue";
import {batchGet} from '@/api/systemsettings'
import {
    HomeOutlined,
    FileOutlined
  } from '@ant-design/icons-vue'

  import {name2label} from '@/utils/helper'


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

    const schemaMap=computed(()=>{
        return schemas.value.reduce((acc, cur) => {
            acc[cur.name] = cur;
            return acc;
        }, {});
    });

    const getSchemaColumns=(name)=>{
       
        return schemaMap.value[name].fields.map((item,index)=>{
            let rs={
                title:name2label(item.name),
                key:item.name,
                dataIndex:item.name,
            
            }
            if (index<3){
                rs.fixed='left'
            }

            if (item.type=='number'){
                rs.width=80
                rs.sorter=(a,b)=>{
                    return a[item.name]-b[item.name]
                }
            }else{
                rs.width=200
            }

            return rs
            
        })

        
    }

    const menuItems=computed(()=>{

        const entities=[];

        for(const entity of settings.value.schemas){
            entities.push({
                key:entity.name,
                label:name2label(entity.name),
                title:entity.name,
                
            });
        }



        return [
            {
                key: 'home',
                icon: () => h(HomeOutlined),
                label: '首页',
                title: '首页',
              },
              {
                key: 'entity',
                icon: () => h(FileOutlined),
                label: '工作表',
                title: '工作表',
                children: entities
              },
        ];
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


    return {settings,getSettings,hasInitialized,schemas,menuItems,schemaMap,getSchemaColumns}

});
