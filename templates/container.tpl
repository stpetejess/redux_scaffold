import React from 'react'
import { connect } from 'react-redux'

import Component from './component'
import { {{.SanitizedName}} } from './action'

class Container extends React.Component {
    render() {
        return <Component {...this.props} />
    }
}
export default connect((s) => {
    return s.{{.SanitizedName}}
}, { {{.SanitizedName}} })(Container)
