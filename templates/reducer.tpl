import { {{.UppercaseSanitizedName}} } from './constant'

export function {{.SanitizedName}}(store={}, action) {
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