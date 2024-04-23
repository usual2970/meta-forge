import service from '@/utils/request'

export const list= (data) => {
  return service({
        url: "/api/v1/data/list",
        method: 'get',
        params: data
    })
}