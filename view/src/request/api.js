import service from '@/request/request'

// @Summary ApiInfo
// @Produce  application/json
export const ApiInfo =  {
    GetApiInfolist : () => {
        return service({
            url: "ApiInfo/GetData",
            method: "POST",
        })
    },
    InsertApiInfo : (data) => {
        return service({
            url: "ApiInfo/Insert",
            method: "POST",
            data: data
        })
    },
    UpdateApiInfo : (data) => {
        return service({
            url: "ApiInfo/Update",
            method: "POST",
            data: data
        })
    },
    DeleteApiInfo : (data) => {
        return service({
            url: "ApiInfo/Delete",
            method: "POST",
            data: data
        })
    }
}


// @Summary ApiJsonInfo
// @Produce  application/json
export const ApiJsonInfo =  {
    GetApiJsonlist : (data) => {
        return service({
            url: "ApiJsonInfo/GetData",
            method: "POST",
            data: data
        })
    },
    InsertApiJsonInfo : (data) => {
        return service({
            url: "ApiJsonInfo/Insert",
            method: "POST",
            data: data
        })
    },
    UpdateApiJsonInfo : (data) => {
        return service({
            url: "ApiJsonInfo/Update",
            method: "POST",
            data: data
        })
    },
    DeleteApiJsonInfo : (data) => {
        return service({
            url: "ApiJsonInfo/Delete",
            method: "POST",
            data: data
        })
    },
    GetJSON : (data) => {
        return service({
            url: "ApiJsonInfo/GetJSON",
            method: "GET",
            params: data
        })
    }
}

// @Summary Domain
// @Produce  application/json
// @Router /Domain [post]
export const Domain =  {
    GetDomainlist : () => {
        return service({
            url: "Domain/GetData",
            method: "POST",
        })
    },
    InsertDomain : (data) => {
        return service({
            url: "Domain/Insert",
            method: "POST",
            data: data
        })
    },
    UpdateDomain : (data) => {
        return service({
            url: "Domain/Update",
            method: "POST",
            data: data
        })
    },
    DeleteDomain : (data) => {
        return service({
            url: "Domain/Delete",
            method: "POST",
            data: data
        })
    }
}