import java.util.*;

public class Planet {

    public String parent;
    public List<String> children;

    public Planet() {
        this.children = new ArrayList<>();
    }

    public String toString() {
        return "parent: " + parent + ", children: " + children + "\n";
    }
}
