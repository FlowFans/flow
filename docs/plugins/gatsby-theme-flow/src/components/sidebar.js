import styled from "@emotion/styled";

import PropTypes from "prop-types";

import React, { Fragment } from "react";

import { Link } from "gatsby";

import { transparentize } from "polished";

import { theme } from "../colors";
import breakpoints from "../utils/breakpoints";
import TrackingLink from "../../../../content/components/tracking-link";

import Logo from "./logo";
import { StyledList, StyledListItem } from "./sidebar-nav";

const Container = styled.aside({
  flexShrink: 0,
  width: 312,
  height: "100vh",
  padding: 24,
  borderRight: `1px solid ${theme.divider}`,
  overflowY: "auto",
  position: "sticky",
  top: 0
});

const ResponsiveContainer = styled(Container)((props) => ({
  [breakpoints.md]: {
    height: "100%",
    backgroundColor: "white",
    boxShadow: `0 0 48px ${transparentize(0.75, "black")}`,
    position: "fixed",
    zIndex: 2,
    opacity: props.open ? 1 : 0,
    visibility: props.open ? "visible" : "hidden",
    transform: props.open ? "none" : "translateX(-100%)",
    transitionProperty: "transform, opacity, visibility",
    transitionDuration: "150ms",
    transitionTimingFunction: "ease-in-out"
  }
}));

const Header = styled.div({
  display: "flex",
  marginBottom: 24
});

const Sidebar = React.forwardRef((props, ref) => {
  const content = (
    <Fragment>
      <Header>
        <TrackingLink href={props.logoLink} eventName="Home_clicked">
          <Logo />
        </TrackingLink>
      </Header>
      <div className={props.className}>
        {props.children || (
          <StyledList>
            <StyledListItem>
              <Link to={"/"}>Home</Link>
            </StyledListItem>
          </StyledList>
        )}
      </div>
    </Fragment>
  );

  if (props.responsive) {
    return (
      <ResponsiveContainer ref={ref} open={props.open}>
        {content}
      </ResponsiveContainer>
    );
  }

  return <Container>{content}</Container>;
});

Sidebar.displayName = "Sidebar";

Sidebar.propTypes = {
  children: PropTypes.node.isRequired,
  open: PropTypes.bool,
  responsive: PropTypes.bool,
  logoLink: PropTypes.string,
  className: PropTypes.string
};

Sidebar.defaultProps = {
  logoLink: "/"
};

export default Sidebar;
