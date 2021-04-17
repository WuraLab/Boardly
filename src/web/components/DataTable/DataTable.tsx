/* eslint-disable react/display-name */
import React, { useEffect, useState } from 'react';
import MaterialTable from 'material-table';
// import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { forwardRef } from 'react';

import AddBox from '@material-ui/icons/AddBox';
import ArrowDownward from '@material-ui/icons/ArrowDownward';
import Check from '@material-ui/icons/Check';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import Clear from '@material-ui/icons/Clear';
import DeleteOutline from '@material-ui/icons/DeleteOutline';
import Edit from '@material-ui/icons/Edit';
import FilterList from '@material-ui/icons/FilterList';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Remove from '@material-ui/icons/Remove';
import SaveAlt from '@material-ui/icons/SaveAlt';
import Search from '@material-ui/icons/Search';
import ViewColumn from '@material-ui/icons/ViewColumn';
// import image from '../../public/logo.svg';
// import { InferGetStaticPropsType } from 'next';

// const useStyles = makeStyles((theme: Theme) => createStyles({}));

const tableIcons = {
    Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
    Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
    Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
    DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
    Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
    Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
    FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
    LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
    NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
    ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
    SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
    ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
    ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />)
};

type User = {
    user_id: string;
    email: string;
    user_name: string;
    password: string;
};

export const getStaticProps = async () => {
    const res = await fetch('https://.../posts');
    const users: User[] = await res.json();

    return {
        revalidate: 10, // revalidate this data , from the backend every 10 seconds
        props: {
            users
        }
    };
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const DataTable: React.FC<{ users: User[] }> = (props) => {
    // const { users } = props;
    const [dataStore, setDataStore] = useState([]);
    // const classes = useStyles();

    useEffect(() => {
        setDataStore([
            {
                profile: 'https://randomuser.me/api/portraits/men/1.jpg',
                name: 'Nextwebb ',
                date: 1987,
                priority: 63
            },
            {
                profile: 'https://randomuser.me/api/portraits/men/2.jpg',
                name: 'Adefemi UI',
                date: 2017,
                priority: 34
            },
            {
                profile: 'https://randomuser.me/api/portraits/men/3.jpg',
                name: 'Seun Mehmet',
                date: 1987,
                priority: 63
            },
            {
                profile: 'https://randomuser.me/api/portraits/men/4.jpg',
                name: 'Robogeek LASU',
                date: 2017,
                priority: 34
            },
            {
                profile: 'https://randomuser.me/api/portraits/men/5.jpg',
                name: 'fawas kareem',
                date: 1987,
                priority: 63
            },
            {
                profile: 'https://randomuser.me/api/portraits/men/6.jpg',
                name: 'Israel  AWS',
                date: 2017,
                priority: 34
            }
        ]);
    }, []);

    return (
        <div style={{ maxWidth: '100%' }}>
            <MaterialTable
                icons={tableIcons}
                title="Employee Datatable"
                columns={[
                    {
                        title: 'Employee Details',
                        field: 'profile',
                        // eslint-disable-next-line react/display-name
                        render: (rowData) => (
                            <img
                                alt="avatar"
                                style={{ height: 36, borderRadius: '50%' }}
                                src={rowData.profile}
                            />
                        ),
                        headerStyle: {
                            color: '#9FA2B4',
                            fontWeight: 'bold'
                        }
                    },
                    {
                        title: 'Employee Name',
                        field: 'name',
                        headerStyle: {
                            color: '#9FA2B4',
                            fontWeight: 'bold'
                        }
                    },
                    {
                        title: 'Date',
                        field: 'date',
                        headerStyle: {
                            color: '#9FA2B4',
                            fontWeight: 'bold'
                        }
                    },
                    {
                        title: 'Priority',
                        field: 'priority',
                        headerStyle: {
                            color: '#9FA2B4',
                            fontWeight: 'bold'
                        }
                    }
                ]}
                data={dataStore}
                options={{
                    search: true
                }}
            />
        </div>
    );
};
