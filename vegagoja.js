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

function vega_version() {
  return vega.version;
}

function vega_lite_version() {
  return vegaLite.version;
}

function vega_render(logf, spec, loadf, cb) {
  const loader = {
    load(name, res) {
      if (res.response != "json" && res.response != "text") {
        logf(["HERERERR", res.response]);
      }
      //logf(["name", name, "res", JSON.stringify(res)]);
      var s = "";
      try {
        s = loadf(name);
      } catch (e) {
        logf(["LOAD ERROR", e]);
      }
      return s;
    },
  };
  try {
    var s = {};
    switch (typeof spec) {
      case "object":
        s = spec;
        break;
      case "string":
        s = JSON.parse(spec);
        break;
      default:
        throw Error("invalid type " + typeof spec);
    }
    var runtime = vega.parse(s);
    var view = new vega.View(runtime, {
      loader: loader,
      logger: logger(logf),
      logLevel: vega.Debug,
    });
    view.toSVG().then(cb);
  } catch (e) {
    logf(["RENDER ERROR"], e);
    throw e;
  } finally {
    if (view) {
      view.finalize();
    }
  }
}

function vega_lite_compile(logf, spec) {
  const s = vegaLite.compile(JSON.parse(spec), {
    logger: logger(logf),
  }).spec;
  return JSON.stringify(s, null, 2);
}

function vega_lite_render(logf, spec, loadf, cb) {
  var s = "";
  try {
    s = vegaLite.compile(JSON.parse(spec), {
      logger: logger(logf),
    }).spec;
  } catch (e) {
    logf(["COMPILE ERROR", e]);
    throw e;
  }
  return vega_render(logf, s, loadf, cb);
}
