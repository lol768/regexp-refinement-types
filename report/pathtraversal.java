public static void writeUserPreference(String username, Preference pref) {
    try {
        String path = "/tmp/prefs/" + username;
        Files.writeString(Paths.get(path), pref.toString());
    } catch(IOException e) {
        e.printStackTrace();
    }
}
