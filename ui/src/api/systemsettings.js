import service from '@/utils/request'



export const get = (data) => {
    return service({
        url: "/api/v1/systemsettings/get",
        method: 'get',
        params: data
    })
}

export const batchGet=(data) => {
    return service({
        url: "/api/v1/systemsettings/batch-get",
        method: 'get',
        params: data
    })
}


export const initialize = (data) => {
    return service({
        url: "/api/v1/systemsettings/initialize",
        method: 'post',
        data
    })
}