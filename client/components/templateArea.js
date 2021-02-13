import { Box, Flex, Grid, Text } from "@chakra-ui/react";
import React from "react";
import EmployeeCards from "./employeeCards";

let cards = [
  {
    title: "Employees",
    text: "60",
  },
  {
    title: "Integrations",
    text: "16",
  },
  {
    title: "Accepted Invites",
    text: "43",
  },
  {
    title: "Pending Invites",
    text: "64",
  },
];
export default function TemplateArea() {
  return (
    <Box px={5}>
      <Grid templateColumns="repeat(4, 1fr)" gap={6}>
        {cards.map((card, i) => (
          <Flex
            key={`card${i}`}
            justify="center"
            align="center"
            bg="#FFFFFF"
            height={150}
            direction="column"
            sx={{ borderRadius: "8px", border: "1px solid #DFE0EB" }}
          >
            <Text fontSize="19px" color="#9FA2B4">
              {card.title}
            </Text>
            <Text fontSize="40px">{card.text}</Text>
          </Flex>
        ))}
      </Grid>

      <Box mt={8}>
        <EmployeeCards />
      </Box>
    </Box>
  );
}
