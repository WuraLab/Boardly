import { Box } from "@chakra-ui/react";
import React from "react";

const EmployeeCards = () => {
  return (
    <Box
      height="300px"
      bg="#FFFFFF"
      p={8}
      sx={{ borderRadius: "8px", border: "1px solid #DFE0EB" }}
    >
      <Box>Employees</Box>
    </Box>
  );
};

export default EmployeeCards;
