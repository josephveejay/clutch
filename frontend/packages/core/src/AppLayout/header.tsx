import React from "react";
import { Link } from "react-router-dom";
import styled from "@emotion/styled";
import { AppBar as MuiAppBar, Box, Grid, Toolbar, Typography } from "@mui/material";

import type { AppConfiguration } from "../AppProvider";
import { FeatureOn, SimpleFeatureFlag } from "../flags";
import { NPSHeader } from "../NPS";

import Logo from "./logo";
import Notifications from "./notifications";
import SearchField from "./search";
import ShortLinker from "./shortLinker";
import { UserInformation } from "./user";

export const APP_BAR_HEIGHT = "64px";

/**
 * Properties used to allow Storybook examples to override featureflag settings
 */
interface HeaderProps extends AppConfiguration {
  /**
   * Will enable the NPS feedback component in the header
   */
  enableNPS?: boolean;
}

const AppBar = styled(MuiAppBar)({
  minWidth: "fit-content",
  background: "linear-gradient(90deg, #38106b 4.58%, #131c5f 89.31%)",
  zIndex: 1201,
  height: APP_BAR_HEIGHT,
});

// Since the AppBar is fixed we need a div to take up its height in order to push other content down.
const ClearAppBar = styled.div({ height: APP_BAR_HEIGHT });

const Title = styled(Typography)({
  margin: "12px 0px 12px 8px",
  fontWeight: "bold",
  fontSize: "30px",
  paddingLeft: "5px",
  color: "rgba(255, 255, 255, 0.87)",
});

const StyledLogo = styled("img")({
  width: "48px",
  height: "48px",
  padding: "1px",
  verticalAlign: "middle",
});

const Header: React.FC<HeaderProps> = ({
  title = "clutch",
  logo = <Logo />,
  enableNPS = false,
}) => {
  const showNotifications = false;

  return (
    <>
      <AppBar position="fixed" elevation={0}>
        <Toolbar>
          <Link to="/">{typeof logo === "string" ? <StyledLogo src={logo} /> : logo}</Link>
          <Title>{title}</Title>
          <Grid container alignItems="center" justifyContent="flex-end">
            <Box>
              <SearchField />
            </Box>
            <SimpleFeatureFlag feature="shortLinks">
              <FeatureOn>
                <ShortLinker />
              </FeatureOn>
            </SimpleFeatureFlag>
            {showNotifications && <Notifications />}
            {enableNPS ? (
              <NPSHeader />
            ) : (
              <SimpleFeatureFlag feature="npsHeader">
                <FeatureOn>
                  <NPSHeader />
                </FeatureOn>
              </SimpleFeatureFlag>
            )}

            <UserInformation />
          </Grid>
        </Toolbar>
      </AppBar>
      <ClearAppBar />
    </>
  );
};

export default Header;
