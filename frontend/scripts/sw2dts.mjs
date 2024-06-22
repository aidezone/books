// const gulp = require("gulp");
// const fetch = require("node-fetch");
// const sw2dts = require("sw2dts");
// const fs = require("fs-extra");
// const path = require("path");
// const gutil = require("gulp-util");

import gutil from "gulp-util";
import path from "path";
import fs from "fs-extra";
import sw2dts from "sw2dts";
import fetch from "node-fetch";
import gulp from "gulp";
import pkg from 'lodash';
const {camelCase} = pkg;
import generateApiEnums from "./api-enums.mjs";
import urlList from "../swagger.config.cjs";

import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// 配置文件
const configFile = "../swagger.config.cjs";
// 读取swagger的文件夹
const filePath = "../.swagger";
// 输出路径
const outputDir = "../src/typings/api/";

/**
 * 由swagger文件生成TS接口文件
 * @param {String} swaggerContent swagger文件内容
 * @param {String} namespace 命名空间
 */
function generateInterface(swaggerContent, namespace) {
  // const camelCase = require("lodash").camelCase;
  const data = JSON.parse(swaggerContent);
  sw2dts
    .convert(data, { namespace })
    .then(dts => {
      const fileName = path.resolve(
        __dirname,
        outputDir + `${camelCase(namespace)}.d.ts`
      );
      const fileContent = dts.replace("namespace ", "namespace Model.");
      fs.outputFile(fileName, fileContent, err => {
        if (err) {
          throw err;
        }

        gutil.log(gutil.colors.green("[√]"), `${namespace} 接口生成完毕`);
      });
    })
    .catch(err => {
      gutil.log(gutil.colors.red(`[${namespace}] ${err}`));
    });
}

const task = async function() {
  if (!fs.pathExistsSync(path.resolve(__dirname, configFile))) {
    gutil.log(
      gutil.colors.red(`Config file 'swagger.config.js' does not exist!`)
    );
    return;
  }

  const files = urlList.map(({ url, namespace, update }) => {
    return {
      url,
      file: `${namespace}.swagger.json`,
      update
    };
  });

  // fs.emptyDirSync(path.resolve(__dirname, filePath));

  const checkStatus = res => {
    if (res.ok) {
      // res.status >= 200 && res.status < 300
      return res;
    } else {
      throw new Error(res.status);
    }
  };

  // 下载swagger文件
  await Promise.all(
    files.map(({ url, file, update }) => {
      if (!update) {
        return;
      }
      fetch(url)
        .then(checkStatus)
        .then(res => res.text())
        .then(body =>
          fs.outputFile(path.resolve(__dirname, filePath, file), body)
        )
        .catch(error =>
          gutil.log(
            gutil.colors.red(`'${url}'`),
            `download failed,`,
            `${error}`
          )
        );
    })
  );

  return gulp
    .src(path.resolve(__dirname, filePath, "*.swagger.json"))
    .on("data", file => {
      // gutil.log('Read %d bytes of data', file.contents.length)
      const namespace = path.basename(file.path).split(".")[0];
      generateInterface(file.contents, namespace);
    })
    .on("end", () => {
      generateApiEnums();
      gutil.log(gutil.colors.green("[√]"), `API 枚举生成完毕`);
    });
}

export default task;