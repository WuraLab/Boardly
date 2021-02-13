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
  Text
} from "@chakra-ui/react";

const Createadmin = () => {
  return (
    <Box>
      <Grid
        templateColumns={"1fr 1fr"}
        sx={{ minHeight: "100vh" }}
        bg="#EEEAEA"
      >
        <Box>
            <Text fontSize="4xl">Create Admin Account</Text>
            <Text fontSize="md">fill out the following information to create an admin account</Text>
          <FormControl id="email" mr={10} ml={15}>
            <FormLabel>Company's Name</FormLabel>
            <Input type="text" placeholder="wuralab" size="sm" />
            <FormLabel>Company's Email</FormLabel>
            <Input type="email" placeholder="user@email.com" size="sm" />
            <FormLabel>Password</FormLabel>
            <Input type="password" placeholder="password" size="sm" />
            <FormLabel>Confirm Password</FormLabel>
            <Input type="password" placeholder="confirm password" size="sm" />
            <Button mt={4} colorScheme="teal" type="submit" size="md">
              Submit
            </Button>
          </FormControl>
        </Box>

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
