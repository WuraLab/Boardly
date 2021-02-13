import React from "react";
import { Box, Text, Grid, Image, Stack, Avatar, Flex } from "@chakra-ui/react";

export default function Navbar() {
  return (
    <Flex justify="space-between" align="center" p={6}>
      <Box>
        <Text sx={{ fontWeight: "bold", fontSize: "24px" }}>Overview</Text>
      </Box>

      <Box>
        <Flex align="center">
          <Text mr={4}>Adefemi Hoodlum</Text>
          <Avatar name="Prosper Otemuyiwa" src="https://bit.ly/prosper-baba" />
        </Flex>
      </Box>
    </Flex>
  );
}
