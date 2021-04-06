import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import { Sidebar } from '../../components/common/SideBar/Sidebar';
import { Main } from '../../components/common/Main/Main';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            // flexGrow: 1,
            // position: 'fixed',
            // height: '100%',
            // width: '100%',
            // margin: '0px',
            // padding: '0px',
            // top: '0px',
            // backgroundColor: '#FFFFFF'
            display: 'flex'
        }
    })
);
const DashboardView: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div className={classes.root}>
                <Grid container spacing={0}>
                    <Sidebar />
                </Grid>
                <Main />
            </div>
        </>
    );
};

export default DashboardView;
