import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
  Box,
  Grid,
  Input,
  Button,
  Center,
  Image,
  Text,
  Flex,
  extendTheme
} from "@chakra-ui/react";

import Link from 'next/link'


const theme = extendTheme({
  textStyles: {
    h1: {
      fontSize: ["48px", "72px"],
      fontWeight: "Bold",
      lineHeight: "110%",
      letterSpacing: "-2%",
      fontFamily: "poppins"
    },
    
  },
})


const Createadmin = () => {
  return (
    <Box>
      <Grid
        templateColumns={"1fr 1fr"}
        sx={{ minHeight: "100vh" }}
        bg="white"
      >
        <Flex bg="white" pt={12} px={6} w={592} align="center" justify="center">
           <Box>
           <Text  textStyle="h1" color="#000000" fontFamily="poppins" fontSize="4xl">Create Admin Account</Text>
            <Text fontSize="md" color="#867979">Fill out the following information to create an admin account</Text>
          <FormControl id="email">
            <FormLabel mt={5}>Company's Name</FormLabel>
            <Input type="text" placeholder="wuralab" size="sm" />
            <FormLabel mt={5}>Company's Email</FormLabel>
            <Input type="email" placeholder="user@email.com" size="sm" />
            <FormLabel mt={5}>Password</FormLabel>
            <Input type="password" placeholder="password" size="sm" />
            <FormLabel mt={5}>Confirm Password</FormLabel>
            <Input type="password" placeholder="confirm password" size="sm" />
            <br />
            <Center>
            <Button width="350px" mt={4} color="white" background= "#22AD80" type="submit" size="md">
            <Link href="/dashboard">
                <a>Create Admin Account</a>
            </Link>
            </Button>
            </Center>
          </FormControl>
          <Center>
          <Link href="/">
            <a href= "/login" >Have an Account ? <span>Login</span></a>
          </Link>
          </Center>
           </Box>
        </Flex>

        <Box bg="white">
          <div className="slanted"> 
          <Center>
                <Image src="/admin-image.png" mt={40}/>
            </Center>
          </div>
           
        </Box>
      </Grid>
    </Box>
  );
};

export default Createadmin;
