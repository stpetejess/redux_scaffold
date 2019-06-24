import { {{.UppercaseSanitizedName}} } from './constant'

export function {{.SanitizedName}}Action() {
    return (d, gs) => {
        d({
            type: {{.UppercaseSanitizedName}},
            data: {
                name: 'your face'
            }
        })
    }
}