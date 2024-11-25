public class LogLevels {
    
    public static String message(String logLine) {
        String[] res = logLine.split(":", 2);
        return res[1].trim();
    }

    public static String logLevel(String logLine) {
        int startIndex = logLine.indexOf('[');
        int endIndex = logLine.indexOf(']', startIndex);
        String res = logLine.substring(startIndex + 1, endIndex);
        return res.toLowerCase();
    }

    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
    }
}
