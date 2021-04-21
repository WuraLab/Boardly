import React from 'react';
import clsx from 'clsx';
import { createStyles, makeStyles, useTheme, Theme } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Badge from '@material-ui/core/Badge';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import MoreIcon from '@material-ui/icons/MoreVert';
import NotificationsIcon from '@material-ui/icons/Notifications';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Overview from '../../../public/overview.svg';
import Employees from '../../../public/employees.svg';
import Integrations from '../../../public/integrations.svg';
import DashboardLogo from '../../../public/dashboardLogo.svg';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            backgroundColor: '#F7F8FC',
            height: '100vh'
        },
        grow: {
            flexGrow: 1
        },
        selected: {},
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
        icon: {
            color: '#fff'
        },
        drawerCloseIcon: {
            color: '#fff'
        },
        onCloseDrawer: {
            [theme.breakpoints.up('sm')]: {
                display: 'none'
            },
            [theme.breakpoints.up('xs')]: {
                display: 'none'
            }
        },
        appBar: {
            zIndex: theme.zIndex.drawer + 1,
            transition: theme.transitions.create(['width', 'margin'], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen
            }),
            [theme.breakpoints.up('sm')]: {
                color: '#000',
                backgroundColor: '#F7F8FC'
                // paddingBottom: '10px'
            },
            [theme.breakpoints.up('xs')]: {
                color: '#000',
                backgroundColor: '#F7F8FC'
                // paddingBottom: '10px'
            }
        },
        customizeToolbar: {
            minHeight: 100,
            border: '0',
            backgroundColor: '#F7F8FC',
            marginTop: '20px'
        },
        title: {
            display: 'none',
            [theme.breakpoints.up('sm')]: {
                display: 'block'
            }
        },
        appBarShift: {
            marginLeft: drawerWidth,
            width: `calc(100% - ${drawerWidth}px)`,
            transition: theme.transitions.create(['width', 'margin'], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen
            })
        },
        sectionDesktop: {
            display: 'none',
            [theme.breakpoints.up('md')]: {
                display: 'flex'
            }
        },
        sectionMobile: {
            display: 'flex',
            [theme.breakpoints.up('md')]: {
                display: 'none'
            }
        },
        menuButton: {
            marginRight: 36
        },
        hide: {
            display: 'none'
        },
        drawer: {
            width: drawerWidth,
            flexShrink: 0,
            whiteSpace: 'nowrap'
        },
        drawerOpen: {
            width: drawerWidth,
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen
            }),

            background: '#16093C',
            color: '#fff'
        },

        drawerClose: {
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen
            }),
            overflowX: 'hidden',
            width: theme.spacing(7) + 1,
            [theme.breakpoints.up('sm')]: {
                width: theme.spacing(9) + 1
            },
            background: '#16093C',
            color: '#fff'
        },
        toolbar: {
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'flex-end',
            padding: theme.spacing(0, 1),
            // necessary for content to be below app bar
            ...theme.mixins.toolbar
            // marginBottom: '25px'
        },
        content: {
            flexGrow: 1,
            padding: theme.spacing(3),
            backgroundColor: '#F7F8FC',
            marginTop: '70px'
        }
    })
);

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export default function Main() {
    const classes = useStyles();
    const theme = useTheme();
    const [open, setOpen] = React.useState(false);

    const handleDrawerOpen = () => {
        setOpen(true);
    };

    const handleDrawerClose = () => {
        setOpen(false);
    };

    const [dashboardTitle, setDashboardTitle] = React.useState('Dashboard');

    const [selectedSidebarItem, setSelectedSidebarItem] = React.useState(false);

    const handleSideBarNaviagtion = (name: string) => {
        setDashboardTitle(name);
        setSelectedSidebarItem(!selectedSidebarItem);
        // console.log(dashboardTitle, selectedSidebarItem);
    };

    // topabar
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const [mobileMoreAnchorEl, setMobileMoreAnchorEl] = React.useState<null | HTMLElement>(null);

    const isMenuOpen = Boolean(anchorEl);
    const isMobileMenuOpen = Boolean(mobileMoreAnchorEl);

    const handleProfileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleMobileMenuClose = () => {
        setMobileMoreAnchorEl(null);
    };

    const handleMenuClose = () => {
        setAnchorEl(null);
        handleMobileMenuClose();
    };

    const handleMobileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
        setMobileMoreAnchorEl(event.currentTarget);
    };

    const menuId = 'primary-search-account-menu';
    const renderMenu = (
        <Menu
            anchorEl={anchorEl}
            anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
            id={menuId}
            keepMounted
            transformOrigin={{ vertical: 'top', horizontal: 'right' }}
            open={isMenuOpen}
            onClose={handleMenuClose}>
            <MenuItem onClick={handleMenuClose}>Profile</MenuItem>
            <MenuItem onClick={handleMenuClose}>My account</MenuItem>
        </Menu>
    );

    const mobileMenuId = 'primary-search-account-menu-mobile';
    const renderMobileMenu = (
        <Menu
            anchorEl={mobileMoreAnchorEl}
            anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
            id={mobileMenuId}
            keepMounted
            transformOrigin={{ vertical: 'top', horizontal: 'right' }}
            open={isMobileMenuOpen}
            onClose={handleMobileMenuClose}>
            <MenuItem>
                <IconButton aria-label="show 11 new notifications" color="inherit">
                    <Badge badgeContent={11} color="secondary">
                        <NotificationsIcon />
                    </Badge>
                </IconButton>
                <p>Notifications</p>
            </MenuItem>
            <MenuItem onClick={handleProfileMenuOpen}>
                <IconButton
                    aria-label="account of current user"
                    aria-controls="primary-search-account-menu"
                    aria-haspopup="true"
                    color="inherit">
                    <AccountCircle />
                </IconButton>
                <p>Profile</p>
            </MenuItem>
        </Menu>
    );

    return (
        <div className={classes.root}>
            <CssBaseline />
            <div>
                <AppBar
                    position="fixed"
                    className={clsx(classes.appBar, {
                        [classes.appBarShift]: open
                    })}
                    style={{
                        background: 'transparent',
                        boxShadow: 'none',
                        backgroundColor: '#F7F8FC'
                    }}>
                    <Toolbar className={classes.customizeToolbar}>
                        {/* <Box mt={4}> */}
                        <IconButton
                            color="inherit"
                            aria-label="open drawer"
                            onClick={handleDrawerOpen}
                            edge="start"
                            className={clsx(classes.menuButton, {
                                [classes.hide]: open
                            })}>
                            <MenuIcon style={{ color: '#16093C' }} />
                        </IconButton>
                        <Typography variant="h6" noWrap className={classes.title}>
                            {dashboardTitle}
                        </Typography>
                        {/* </Box> */}

                        <div className={classes.grow} />
                        <div className={classes.sectionDesktop}>
                            <IconButton aria-label="show 17 new notifications" color="inherit">
                                <Badge badgeContent={17} color="secondary">
                                    <NotificationsIcon />
                                </Badge>
                            </IconButton>

                            <IconButton
                                edge="end"
                                aria-label="account of current user"
                                aria-controls={menuId}
                                aria-haspopup="true"
                                onClick={handleProfileMenuOpen}
                                color="inherit">
                                <Box p={1} borderLeft={1} borderColor={'#252733'} color={'#252733'}>
                                    <Typography className={classes.title} variant="h6" noWrap>
                                        Adefemi hoodlum
                                    </Typography>
                                </Box>

                                <AccountCircle />
                            </IconButton>
                        </div>
                        <div className={classes.sectionMobile}>
                            <IconButton
                                aria-label="show more"
                                aria-controls={mobileMenuId}
                                aria-haspopup="true"
                                onClick={handleMobileMenuOpen}
                                color="inherit">
                                <MoreIcon />
                            </IconButton>
                        </div>
                    </Toolbar>
                </AppBar>
                {renderMobileMenu}
                {renderMenu}
            </div>
            <Drawer
                variant="permanent"
                className={clsx(classes.drawer, {
                    [classes.drawerOpen]: open,
                    [classes.drawerClose]: !open
                })}
                classes={{
                    paper: clsx({
                        [classes.drawerOpen]: open,
                        [classes.drawerClose]: !open
                    })
                }}>
                <div className={classes.toolbar}>
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
                    <IconButton onClick={handleDrawerClose} className={classes.drawerCloseIcon}>
                        {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
                    </IconButton>
                </div>
                <Divider />
                <List>
                    <ListItem
                        button
                        classes={{ root: classes.listItem, selected: classes.selected }}
                        selected={selectedSidebarItem}
                        onClick={() => handleSideBarNaviagtion('Overview')}>
                        <ListItemIcon className={classes.icon}>
                            <img alt={'Overview'} src={Overview} />
                        </ListItemIcon>
                        <ListItemText primary={'Overview'} />
                    </ListItem>
                    <ListItem
                        button
                        classes={{ root: classes.listItem, selected: classes.selected }}
                        // selected={selectedSidebarItem}
                        onClick={() => handleSideBarNaviagtion('Employees')}>
                        <ListItemIcon className={classes.icon}>
                            <img alt={'Employees'} src={Employees} />
                        </ListItemIcon>
                        <ListItemText primary={'Employees'} />
                    </ListItem>
                    <ListItem
                        button
                        classes={{ root: classes.listItem, selected: classes.selected }}
                        // selected={selectedSidebarItem}
                        onClick={() => handleSideBarNaviagtion('Integration')}>
                        <ListItemIcon className={classes.icon}>
                            <img alt={'Integration'} src={Integrations} />
                        </ListItemIcon>
                        <ListItemText primary={'Integration'} />
                    </ListItem>
                </List>
            </Drawer>
            <main className={classes.content}>
                <div className={classes.toolbar} />
                {/* <Typography variant="h6" noWrap className={classes.title}> */}
                {/* {dashboardTitle} */}
                {/* </Typography> */}
            </main>
        </div>
    );
}
