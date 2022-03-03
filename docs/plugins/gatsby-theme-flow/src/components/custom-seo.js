import PropTypes from "prop-types";

import React from "react";

import { withPrefix } from "gatsby";

import Seo from "./seo";

export default function CustomSEO({ image, baseUrl, twitterHandle, ...props }) {
  return (
    <Seo {...props} twitterCard="summary_large_image">
      <meta property="og:image" content={image} />
      {baseUrl && <meta name="twitter:image" content={image} />}
      {twitterHandle && (
        <meta name="twitter:site" content={`@${twitterHandle}`} />
      )}
    </Seo>
  );
}

CustomSEO.propTypes = {
  baseUrl: PropTypes.string,
  image: PropTypes.string.isRequired,
  twitterHandle: PropTypes.string
};
