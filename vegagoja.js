Date.prototype.setYear = Date.prototype.setFullYear;

const protocol_re = /^(data:|([A-Za-z]+:)?\/\/)/;
const allowed_re =
  /^(?:(?:(?:f|ht)tps?|mailto|tel|callto|cid|xmpp|file|data):|[^a-z]|[a-z+.\-]+(?:[^a-z+.\-:]|$))/i;
const whitespace_re =
  /[\u0000-\u0020\u00A0\u1680\u180E\u2000-\u2029\u205f\u3000]/g;

function logger(logf) {
  return {
    level(_) {},
    error() {
      logf(["ERROR", ...arguments]);
    },
    warn() {
      logf(["WARN", ...arguments]);
    },
    info() {
      logf(["INFO", ...arguments]);
    },
    debug() {
      logf(["DEBUG", ...arguments]);
    },
  };
}

function parse(spec) {
  if (typeof spec == "object") {
    return spec;
  }
  return JSON.parse(spec);
}

function version() {
  return [vega.version, vegaLite.version];
}

function compile(logf, spec) {
  let res;
  try {
    res = vegaLite.compile(parse(spec), {
      logger: logger(logf),
    });
  } catch (e) {
    logf(["COMPILE ERROR", e]);
    throw e;
  }
  return JSON.stringify(res);
}

function loader(logf, loadf) {
  return {
    async load(name, res) {
      let s;
      try {
        s = loadf(name, res.response);
      } catch (e) {
        logf(["LOAD ERROR", e]);
        throw e;
      }
      return s;
    },
    // stripped down version of vega's sanitize
    async sanitize(uri, options) {
      const result = {
        href: null,
      };
      let base;
      let isAllowed = allowed_re.test(uri.replace(whitespace_re, ""));
      if (uri == null || typeof uri !== "string" || !isAllowed) {
        throw new Error("Sanitize failure, invalid URI: " + uri);
      }
      const hasProtocol = protocol_re.test(uri);
      if ((base = options.baseURL) && !hasProtocol) {
        if (!uri.startsWith("/") && !base.endsWith("/")) {
          uri = "/" + uri;
        }
        uri = base + uri;
      }
      result.href = uri;
      if (options.target) {
        result.target = options.target + "";
      }
      if (options.rel) {
        result.rel = options.rel + "";
      }
      if (options.context === "image" && options.crossOrigin) {
        result.crossOrigin = options.crossOrigin + "";
      }
      return result;
    },
  };
}

async function render(logf, spec, loadf) {
  let runtime = vega.parse(parse(spec));
  let view = new vega.View(runtime, {
    logLevel: vega.Debug,
    logger: logger(logf),
    loader: loader(logf, loadf),
  });
  let result = await view.toSVG().finally(() => {
    view.finalize();
  });
  return result;
}
