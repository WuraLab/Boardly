import React, { useState } from "react";
import { Box, Text, Grid, Image, Stack } from "@chakra-ui/react";
import Sidebar from "../components/sidebar";
import Navbar from "../components/navbar";

const Dashboard = () => {
  return (
    <Box>
      <Grid
        templateColumns={"auto 1fr"}
        templateRows={"auto 1fr"}
        sx={{ minHeight: "100vh" }}
        bg="#F7F8FC"
        templateAreas="'sidebar navbar' 'sidebar body'"
      >
        <Box sx={{ gridArea: "sidebar", height: "100%" }}>
          <Sidebar />
        </Box>

        <Box sx={{ gridArea: "navbar" }}>
          <Navbar />
        </Box>
        <Box sx={{ gridArea: "body" }}> body</Box>
      </Grid>
    </Box>
  );
};

export default Dashboard;
