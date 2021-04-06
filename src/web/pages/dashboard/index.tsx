import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import { Sidebar } from '../../components/common/SideBar/Sidebar';

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
        rightGrid: {
            position: 'relative',
            backgroundColor: '#FFFFFF',
            height: '100vh',
            border: '1px solid blue'
        },
        leftGrid: {
            height: '100%',
            backgroundColor: '#16093C',
            // border: '1px solid red',
            // justifyContent: 'center',
            // alignContent: 'center',
            padding: 0,
            margin: 0
        }
    })
);
const DashboardView: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div className={classes.root}>
                <Grid container spacing={0}>
                    <Grid item xs={12} sm={3} className={classes.leftGrid}>
                        <Sidebar />
                    </Grid>

                    <Grid item xs={12} sm={9} className={classes.rightGrid}>
                        <h1>main</h1>
                    </Grid>
                </Grid>
            </div>
        </>
    );
};

export default DashboardView;
