import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'

import styles from './styles'

const Component = ({ classes, {{.SanitizedName}} }) => {
    return (
        <div className={classes.root}>
            { `Hello ${ {{.SanitizedName}} }` }
        </div>
    )
}

Component.PropTypes = {
    classes: PropTypes.object.isRequired,
    {{.SanitizedName}}: PropTypes.func.isRequired
}

export default withStyles(styles)(Component)