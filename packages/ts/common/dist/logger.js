"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Logger = exports.LogLevel = void 0;
var LogLevel;
(function (LogLevel) {
    LogLevel["DEBUG"] = "DEBUG";
    LogLevel["INFO"] = "INFO";
    LogLevel["WARN"] = "WARN";
    LogLevel["ERROR"] = "ERROR";
})(LogLevel || (exports.LogLevel = LogLevel = {}));
class Logger {
    constructor(context) {
        this.context = context;
    }
    formatMessage(level, message) {
        const timestamp = new Date().toISOString();
        return `[${timestamp}] [${level}] [${this.context}] ${message}`;
    }
    debug(message) {
        console.debug(this.formatMessage(LogLevel.DEBUG, message));
    }
    info(message) {
        console.info(this.formatMessage(LogLevel.INFO, message));
    }
    warn(message) {
        console.warn(this.formatMessage(LogLevel.WARN, message));
    }
    error(message, error) {
        const fullMessage = error
            ? `${message} - ${error.message}\n${error.stack}`
            : message;
        console.error(this.formatMessage(LogLevel.ERROR, fullMessage));
    }
}
exports.Logger = Logger;
