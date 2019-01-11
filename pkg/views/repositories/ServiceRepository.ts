import * as React from 'react'


export interface GetConfigurationResponseBody {
    distributorAddress: string
    assetName: string
    issuerAddress: string
    candidates: {name: string, address: string}[]
}

export interface ServiceRepository {
    getConfiguration(): Promise<GetConfigurationResponseBody>
}

export class DefaultServiceRepository implements ServiceRepository{
    async getConfiguration(): Promise<GetConfigurationResponseBody> {
            const resp = await fetch("/configuration")
            const result = await resp.json()
            return {
                assetName: result.assetName,
                candidates: result.candidates,
                distributorAddress: result.distributorAddress,
                issuerAddress: result.issuerAddress
            }
    }
}

export const ServiceContext = React.createContext(new DefaultServiceRepository())
