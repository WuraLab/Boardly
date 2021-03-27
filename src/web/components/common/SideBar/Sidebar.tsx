import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { Grid, Box } from '@material-ui/core';
import Logo from '../../../public/logo.svg';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
            position: 'fixed',
            height: '100%',
            width: '100%',
            margin: '0px',
            padding: '0px',
            top: '0px',
            backgroundColor: '#FFFFFF'
        },

        navItem: {
            backgroundColor: '#FFFFFF',
            border: '1px solid red',
            padding: 0,
            margin: 0
        }
    })
);

export const Sidebar: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div className={classes.root}>
                <Grid container spacing={3}>
                    <Grid item xs={12} sm={12}>
                        <Box>
                            <img
                                alt="forgot password"
                                style={{
                                    height: '100px',
                                    maxWidth: '100px',
                                    marginLeft: '100px'
                                }}
                                src={Logo}
                            />
                        </Box>
                    </Grid>

                    <Grid item xs={12} sm={12} className={classes.navItem}>
                        <h1>Overview</h1>
                    </Grid>

                    <Grid item xs={12} sm={12} className={classes.navItem}>
                        <h1>Employees</h1>
                    </Grid>
                </Grid>
            </div>
        </>
    );
};
