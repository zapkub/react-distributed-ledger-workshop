import * as React from 'react'
import {HorizonContext} from "../repositories/StellarRepository";
import {ServiceContext} from "../repositories/ServiceRepository";

export enum AddressValidateStatus {
    NONE,
    INVALID,
    VALID
}

export function useAddressValidation() {
    const [address, setAddress] = React.useState("");
    const [isValidate, setValidate ] = React.useState(AddressValidateStatus.NONE)
    const [balance, setBalance] = React.useState(0)

    const horizon = React.useContext(HorizonContext);
    const service = React.useContext(ServiceContext)


    return {
        isValidate,
        address,
        balance,
        setAddress: (address: string) => {
            setAddress(address);
            setValidate(AddressValidateStatus.NONE)
        },
        doValidateAddress: async () => {

            // Validate address if it can
            // vote for BlackPink by checking
            // balance information

            const config = await service.getConfiguration()

            try {
                const account = await horizon.loadAccount(address)
                const voteAsset = account.balances.find((asset) => {
                    if (asset.asset_type == "credit_alphanum4") {
                        return asset.asset_issuer == config.issuerAddress && asset.asset_code == config.assetName
                    } else {
                        return false
                    }
                })

                if ( voteAsset ) {
                    setValidate(AddressValidateStatus.VALID)
                    setBalance(parseInt(voteAsset.balance))
                } else {
                    console.error(voteAsset)
                    setValidate(AddressValidateStatus.INVALID)
                }
            } catch(e) {
                console.error(e)
                setValidate(AddressValidateStatus.INVALID)
            }

        }
    }
}