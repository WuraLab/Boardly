import React, { useState } from "react";
import { Box, Text, Grid, Image, Stack } from "@chakra-ui/react";

const Sidebar = () => {
  const [activeTab, setActiveTab] = useState(0);

  const tabs = [
    {
      text: "overview",
      icon: "",
    },
    {
      text: "Employees",
      icon: "",
    },
    {
      text: "Integrations",
      icon: "",
    },
  ];

  return (
    <Box
      sx={{
        height: "100%",
        minWidth: "225px",
        color: "#ffffff",
      }}
      bg="#16093C"
    >
      {/* logo area */}

      <Stack p={7} direction={["column", "row"]} align="center" spacing="10px">
        <Box>
          <Image src="/logo.png" />
        </Box>
        <Text sx={{ fontWeight: "bold", fontSize: "19px" }}>Boardly</Text>
      </Stack>

      {/* tabs */}
      <Box mt={14}>
        {tabs.map((tab, i) => (
          <Grid templateColumns="auto 1fr" sx={{ cursor: "pointer" }}>
            <Box
              sx={{
                width: "6px",
                height: "100%",
                bg: activeTab === i ? "#DDE2FF" : null,
              }}
            ></Box>
            <Box p={7} sx={{ bg: activeTab === i ? "#105256" : null }}>
              {/* todo: icon */}
              <Text>{tab.text}</Text>
            </Box>
          </Grid>
        ))}
      </Box>
    </Box>
  );
};

export default Sidebar;
