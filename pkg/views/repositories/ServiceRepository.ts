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
    private configURL = "/configuration"
    async getConfiguration(): Promise<GetConfigurationResponseBody> {
        throw new Error("Implement me")
    }
}

export const ServiceContext = React.createContext(new DefaultServiceRepository())
