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
  Flex
} from "@chakra-ui/react";

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
           <Text fontSize="4xl">Create Admin Account</Text>
            <Text fontSize="md">fill out the following information to create an admin account</Text>
          <FormControl id="email">
            <FormLabel mt={5}>Company's Name</FormLabel>
            <Input type="text" placeholder="wuralab" size="sm" />
            <FormLabel mt={5}>Company's Email</FormLabel>
            <Input type="email" placeholder="user@email.com" size="sm" />
            <FormLabel mt={5}>Password</FormLabel>
            <Input type="password" placeholder="password" size="sm" />
            <FormLabel mt={5}>Confirm Password</FormLabel>
            <Input type="password" placeholder="confirm password" size="sm" />
            <Button mt={4} colorScheme="teal" type="submit" size="md">
              Submit
            </Button>
          </FormControl>
           </Box>
        </Flex>

        <Box bg="#7742F6">

            <Center>
                <Image src="/admin-image.png" mt={40}/>
            </Center>
        </Box>
      </Grid>
    </Box>
  );
};

export default Createadmin;
