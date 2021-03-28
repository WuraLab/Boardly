import { createMuiTheme } from '@material-ui/core/styles';
import indigo from '@material-ui/core/colors/indigo';
import pink from '@material-ui/core/colors/pink';

export default createMuiTheme({
    palette: {
        type: 'light',
        primary: {
            main: '#7e86f6'
        },
        secondary: {
            main: '#22AD80'
        }
    },
    typography: {
        fontFamily: 'Poppins',
        fontSize: 13
    }
});
