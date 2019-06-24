import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'

import styles from './styles'

const Component = ({ classes, name, {{.SanitizedName}} }) => {
    return (
        <div className={classes.root}>
            { `Hello ${ name }` }
        </div>
    )
}

Component.PropTypes = {
    classes: PropTypes.object.isRequired,
    name: PropTypes.string.isRequired,
    {{.SanitizedName}}Action: PropTypes.func.isRequired
}

export default withStyles(styles)(Component)