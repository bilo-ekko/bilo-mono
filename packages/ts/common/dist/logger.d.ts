export declare enum LogLevel {
    DEBUG = "DEBUG",
    INFO = "INFO",
    WARN = "WARN",
    ERROR = "ERROR"
}
export declare class Logger {
    private context;
    constructor(context: string);
    private formatMessage;
    debug(message: string): void;
    info(message: string): void;
    warn(message: string): void;
    error(message: string, error?: Error): void;
}
