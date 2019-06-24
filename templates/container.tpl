import React from 'react'
import { connect } from 'react-redux'

import Component from './component'
import { {{.SanitizedName}}Action } from './action'

class Container extends React.Component {
    render() {
        return <Component {...this.props} />
    }
}
export default connect((s) => {
    return s.{{.LowerSanitizedName}}
}, { {{.SanitizedName}}Action })(Container)
