function vega_version() {
  return vega.version;
}

function vega_lite_version() {
  return vegaLite.version;
}

function vega_render(logf, spec, loadf, cb) {
  const logger = {
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
      logger: logger,
      logLevel: vega.Debug,
    });
    view.toSVG().then(cb);
  } catch (e) {
    throw e;
  } finally {
    if (view) {
      view.finalize();
    }
  }
}

function vega_lite_render(logf, spec, loadf, cb) {
  var s = "";
  try {
    var v = JSON.parse(spec);
    s = vegaLite.compile(v).spec;
  } catch (e) {
    logf(["COMPILE ERROR", e]);
    throw e;
  }
  logf(["COMPILED", JSON.stringify(s, null, 2)]);
  return vega_render(logf, s, loadf, cb);
}
