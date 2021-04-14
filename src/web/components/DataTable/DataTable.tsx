import React, { FC, useState } from 'react';
import MaterialTable from 'material-table';
import DeleteIcon from '@material-ui/icons/Delete';
import SearchIcon from '@material-ui/icons/Search';
import SaveIcon from '@material-ui/icons/Save';
import { Button } from '@material-ui/core';

import { useStyles } from './material-table-wrapper.styles';

interface Column {
    name: string;
    occupation: string;
    age: number;
}

export const DataTable: FC = () => {
    // const classes = useStyles();

    return (
        <div style={{ maxWidth: '100%' }}>
            <MaterialTable
                title="Remote Data Preview"
                columns={[
                    {
                        title: 'Avatar',
                        field: 'avatar',
                        render: (rowData) => (
                            <img style={{ height: 36, borderRadius: '50%' }} src={rowData.avatar} />
                        )
                    },
                    { title: 'Id', field: 'id' },
                    { title: 'First Name', field: 'first_name' },
                    { title: 'Last Name', field: 'last_name' }
                ]}
                data={(query) =>
                    new Promise((resolve, reject) => {
                        let url = 'https://reqres.in/api/users?';
                        url += 'per_page=' + query.pageSize;
                        url += '&page=' + (query.page + 1);
                        fetch(url)
                            .then((response) => response.json())
                            .then((result) => {
                                resolve({
                                    data: result.data,
                                    page: result.page - 1,
                                    totalCount: result.total
                                });
                            });
                    })
                }
            />
        </div>
    );
};

// export const DataTable: FC = () => {
//     // const classes = useStyles();
//     const [dataStore, setDataStore] = useState([{ name: 'Jon', job: 'Software Dev', age: 29 }]);

//     return (
//         <div style={{ maxWidth: '100%' }}>
//             <MaterialTable
//                 columns={[
//                     {
//                         title: 'Name',
//                         field: 'name'
//                     },
//                     {
//                         title: 'Occupation',
//                         field: 'job'
//                     },
//                     {
//                         title: 'Age',
//                         field: 'age',
//                         type: 'numeric'
//                     }
//                 ]}
//                 data={dataStore}
//                 title="Material-Table Demo"
//                 icons={{
//                     Clear: ((() => <DeleteIcon />) as unknown) as React.ForwardRefExoticComponent<
//                         React.RefAttributes<SVGSVGElement>
//                     >,
//                     Search: ((() => <SearchIcon />) as unknown) as React.ForwardRefExoticComponent<
//                         React.RefAttributes<SVGSVGElement>
//                     >,
//                     ResetSearch: ((() => (
//                         <DeleteIcon />
//                     )) as unknown) as React.ForwardRefExoticComponent<
//                         React.RefAttributes<SVGSVGElement>
//                     >
//                 }}
//                 actions={[
//                     {
//                         icon: () => <SaveIcon />,
//                         tooltip: 'Save User',
//                         onClick: (event, rowData) =>
//                             alert('You saved ' + ((rowData as unknown) as Column).name)
//                     }
//                 ]}
//                 components={{
//                     // eslint-disable-next-line react/display-name
//                     Action: (props) => (
//                         <Button
//                             onClick={(event) => props.action.onClick(event, props.data)}
//                             color="primary"
//                             variant="text"
//                             style={{ textTransform: 'none' }}
//                             size="small">
//                             Save
//                         </Button>
//                     )
//                 }}
//                 options={{
//                     headerStyle: {
//                         backgroundColor: '#01579b',
//                         color: '#FFF'
//                     },
//                     search: true
//                 }}
//             />
//         </div>
//     );
// };
