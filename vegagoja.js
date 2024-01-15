function vega_version() {
  return vega.version;
}

function vega_lite_version() {
  return "";
  //return vegaLite.version;
}

function render(logf, loadf, cb, spec, data) {
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
    var s = JSON.parse(spec);
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
