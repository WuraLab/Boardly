import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import Main from '../../components/Main/Main';
import CssBaseline from '@material-ui/core/CssBaseline';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            backgroundColor: '#F7F8FC',
            // backgroundColor: 'red',
            border: '1px solid black'
        }
    })
);
const DashboardView: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div className={classes.root}>
                {/* <Grid container spacing={0}> */}
                <CssBaseline />
                <Main />
                {/* </Grid> */}
            </div>
        </>
    );
};

export default DashboardView;
