import { {{.UppercaseSanitizedName}} } from './constant'

export default function {{.LowerSanitizedName}}(store={}, action) {
    switch(action.type) {
        case {{.UppercaseSanitizedName}}: {
            return {
                ...store,
                ...action.data
            }
        }
        default: return store
    }
}