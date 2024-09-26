import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;

public class Day06 {

    public static Map<String, Planet> planets = new HashMap<>();

    public static void main(String[] args) {
        List<String> lines = new ArrayList<String>();
        String filePath = "input.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            String line;
            while ((line = br.readLine()) != null) {
                lines.add(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

        for (int i = 0; i < lines.size(); i++) {
            String[] parts = lines.get(i).split("\\)");
            String from = parts[0];
            String to = parts[1];

            if (planets.get(from) == null) {
                planets.put(from, new Planet());
            }
            if (planets.get(to) == null) {
                planets.put(to, new Planet());
            }
            planets.get(from).children.add(to);
            planets.get(to).parent = from;
        }
        // System.out.println(planets);

        // count orbits
        System.out.println(getNumOrbits(planets.get("COM"), 1));

        List<String> myPathBack = getPathBack(planets.get("YOU"));
        List<String> santasPathBack = getPathBack(planets.get("SAN"));

        int numCommon = 0;
        for (
            ;
            myPathBack.get(numCommon) == santasPathBack.get(numCommon);
            numCommon++
        ) {}
        numCommon++;
        System.out.println(
            myPathBack.size() - numCommon + santasPathBack.size() - numCommon
        );
    }

    public static int getNumOrbits(Planet parent, int depth) {
        int orbits = 0;
        for (int i = 0; i < parent.children.size(); i++) {
            orbits += depth;
            Planet child = planets.get(parent.children.get(i));
            orbits += getNumOrbits(child, depth + 1);
        }
        return orbits;
    }

    public static List<String> getPathBack(Planet p) {
        List<String> path = new ArrayList<>();
        while (p.parent != null) {
            path.add(0, p.parent);
            p = planets.get(p.parent);
        }
        return path;
    }
}
