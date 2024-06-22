// const requireContext = require("require-context");
// const apiGenerator = require("../src/fetch/apiGenerator");
// const path = require("path");
// const fs = require("fs-extra");
import fs from "fs-extra";
import path from "path";
import apiGenerator from "../src/fetch/apiGenerator/index.mjs"
import requireContext from "require-context";

import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// 读取swagger的文件夹
const filePath = "../.swagger";
const outputDir = "../src/fetch/api";

function removeLineBreaks(str) {
  return str && str.replace(/[\r\n]/g, "");
}

function getSchemaType(ref) {
  const result = ref.replace("#/definitions/", "").replace(/[.]/g, "");
  return result.substring(0, 1).toUpperCase() + result.substring(1);
}

function getParamType(param, ns) {
  return (
    param?.type ||
    (param.schema
      ? param.schema.type ||
        "Model." + ns + "." + getSchemaType(param.schema.$ref)
      : "")
  );
}

function getEnumContent({ namespace, operations }) {
  const genParameters = operation => {
    return (operation.parameters || [])
      .map(p => {
        console.log("operation.parameters.map",p)
        const type = getParamType(p, namespace);
        const description = removeLineBreaks(p.description);
        const isRequired = p.required ? "(required)" : "";

        return `
     * @param {${type}} ${p.name} - [${p.in}] ${description} ${isRequired}`;
      })
      .join("");
  };

  return `// ${namespace}
  export enum ${namespace} {${operations
    .map(o => {
      const description = removeLineBreaks(o.description || o.summary);
      const parameters = genParameters(o);
      const type = getParamType(o.responses["200"], namespace);
      const method = o.method.toUpperCase();

      return `
    /**
     * ${description}
     * ${parameters}
     * @return {${type}} ${type}
     * @summary ${method} ${o.path}
     */
    ${o.operationId} = '${namespace}:${o.operationId}',`;
    })
    .join("\n")}
  }
`;
}

/**
 * 生成待处理的swagger内容
 * @param { swagger } swagger swagger内容
 */
function generateEnumContent(swagger) {
  const namespace = swagger.namespace;
  const operations = [];

  const paths = swagger.path;
  Object.keys(paths).forEach(key => {
    const path = paths[key];
    Object.keys(path).forEach(method => {
      const operation = path[method];
      operation.method = method;
      operation.path = key;
      operations.push(operation);
    });
  });

  return {
    namespace,
    operations
  };
}

export default function() {
  const context = requireContext(
    path.resolve(__dirname, filePath),
    true,
    /\.swagger\.json$/i
  );

  const swaggers = apiGenerator.collectSwaggerInfo(context);
  let enumContent = "";

  swaggers.forEach(swagger => {
    const content = generateEnumContent(swagger);
    enumContent += `\n  ${getEnumContent(content)}`;
  });

  const outputPath = path.resolve(__dirname, outputDir, "index.ts");

  const apiEnums = `/* This file was auto-generated. Don't modify this manually. */\n
export namespace API {${enumContent}}
`;
  fs.outputFile(outputPath, apiEnums);
};