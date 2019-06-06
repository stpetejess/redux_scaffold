import { {{.UppercaseSanitizedName}} } from './constant'

export function {{.SanitizedName}}() {
    return (d, gs) => {
        d({
            type: {{.UppercaseSanitizedName}},
            data: {
                name: 'your face'
            }
        })
    }
}