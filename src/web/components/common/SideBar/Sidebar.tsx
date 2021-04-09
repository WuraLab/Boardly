/* eslint-disable @typescript-eslint/explicit-function-return-type */
import React from 'react';
import TopBar from '../../common/TopBar/Topbar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Typography from '@material-ui/core/Typography';
import AppBar from '@material-ui/core/AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Divider from '@material-ui/core/Divider';
import Drawer from '@material-ui/core/Drawer';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Hidden from '@material-ui/core/Hidden';
import Overview from '../../../public/overview.svg';
import Employees from '../../../public/employees.svg';
import Integrations from '../../../public/integrations.svg';
import DashboardLogo from '../../../public/dashboardLogo.svg';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import { makeStyles, useTheme, Theme, createStyles } from '@material-ui/core/styles';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex'
        },
        listItemText: {
            fontSize: '2.0em' //Insert your required size
        },
        listItem: {
            '&$selected': {
                backgroundColor: '#105256',
                '&:hover': {
                    backgroundColor: '#105256'
                }
            }
        },
        topNavigation: {
            backgroundColor: 'blue',
            float: 'right',
            width: '400px',
            marginRight: 0,
            paddingRight: 0
        },
        selected: {},
        icon: {
            color: '#fff'
        },
        drawer: {
            [theme.breakpoints.up('sm')]: {
                width: drawerWidth,
                flexShrink: 0
            }
        },
        appBar: {
            [theme.breakpoints.up('sm')]: {
                width: `calc(100% - ${drawerWidth}px)`,
                marginLeft: drawerWidth,
                background: '#fff',
                color: '#000'
            }
        },
        menuButton: {
            marginRight: theme.spacing(2),
            [theme.breakpoints.up('sm')]: {
                display: 'none'
            }
        },
        // necessary for content to be below app bar
        toolbar: theme.mixins.toolbar,
        drawerPaper: {
            width: drawerWidth,
            background: '#16093C',
            color: '#fff'
        },
        content: {
            flexGrow: 1,
            padding: theme.spacing(3)
        }
    })
);

interface Props {
    window?: () => Window;
}

export const Sidebar: React.FC = (props: Props) => {
    const { window } = props;
    const classes = useStyles();
    const theme = useTheme();
    const [mobileOpen, setMobileOpen] = React.useState(false);
    const [dashboardTitle, setDashboardTitle] = React.useState('Dashboard');

    const [selectedSidebarItem, setSelectedSidebarItem] = React.useState(false);

    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen);
    };

    const handleSideBarNaviagtion = (name: string) => {
        setDashboardTitle(name);
        setSelectedSidebarItem(!selectedSidebarItem);
        // console.log(dashboardTitle, selectedSidebarItem);
    };

    const drawer = (
        <div>
            <Grid>
                <Box
                    className={classes.toolbar}
                    display="flex"
                    justifyContent="center"
                    mt={2}
                    mb={4}>
                    <ListItem>
                        <ListItemIcon className={classes.icon}>
                            <img alt="logo" src={DashboardLogo} />
                        </ListItemIcon>
                        <ListItemText
                            primary="Boardly"
                            classes={{ primary: classes.listItemText }}
                        />
                    </ListItem>
                </Box>
            </Grid>

            <Divider />
            <List>
                {[
                    { name: 'Overview', src: Overview, selected: selectedSidebarItem },
                    { name: 'Employees', src: Employees },
                    { name: 'Integration', src: Integrations }
                ].map((item, index) => (
                    <ListItem
                        button
                        key={index}
                        classes={{ root: classes.listItem, selected: classes.selected }}
                        selected={item.selected}
                        onClick={() => handleSideBarNaviagtion(item.name)}>
                        <ListItemIcon className={classes.icon}>
                            <img alt={item.name} src={item.src} />
                        </ListItemIcon>
                        <ListItemText primary={item.name} />
                    </ListItem>
                ))}
            </List>
            <Divider />
        </div>
    );

    const container = window !== undefined ? () => window().document.body : undefined;

    return (
        <>
            <div>
                <CssBaseline />

                {/* top bar */}
                {/* <AppBar position="fixed" className={classes.appBar}>
                    <TopBar dashboardTitle={dashboardTitle} />
                </AppBar> */}
                <AppBar position="fixed" className={classes.appBar}>
                    <Toolbar>
                        <IconButton
                            color="inherit"
                            aria-label="open drawer"
                            edge="start"
                            onClick={handleDrawerToggle}
                            className={classes.menuButton}>
                            <MenuIcon />
                        </IconButton>
                        <Typography variant="h6" noWrap className={classes.title}>
                            {dashboardTitle}
                        </Typography>
                    </Toolbar>
                </AppBar>

                {/* side navigation */}
                <nav className={classes.drawer} aria-label="mailbox folders">
                    {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
                    <Hidden smUp implementation="css">
                        <Drawer
                            container={container}
                            variant="temporary"
                            anchor={theme.direction === 'rtl' ? 'right' : 'left'}
                            open={mobileOpen}
                            onClose={handleDrawerToggle}
                            classes={{
                                paper: classes.drawerPaper
                            }}
                            ModalProps={{
                                keepMounted: true // Better open performance on mobile.
                            }}>
                            {drawer}
                        </Drawer>
                    </Hidden>
                    <Hidden xsDown implementation="css">
                        <Drawer
                            classes={{
                                paper: classes.drawerPaper
                            }}
                            variant="permanent"
                            open>
                            {drawer}
                        </Drawer>
                    </Hidden>
                </nav>
            </div>
        </>
    );
};
