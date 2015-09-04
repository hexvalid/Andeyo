/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package telegram;

import java.util.Arrays;

/**
 *
 * @author erkanmdr
 */
public class Cevirici {

    public static String[] latin_alfabesi = {"a", "b", "c", "ç", "d", "e", "f", "g", "ğ", "h", "ı", "i", "j", "k", "l", "m", "n", "o", "ö", "p", "r", "s", "ş", "t", "u", "ü", "v", "y", "z"};

    public static void main(String[] args) {

    }

    public static String andeyoConverter(String latin) {
        String kucuklatin = latin.toLowerCase();
        kucuklatin = kucuklatin.replaceAll("a", "🎄")
                .replaceAll("b", "🔩")
                .replaceAll("c", "🌜")
                .replaceAll("ç", "💭")
                .replaceAll("d", "🌛")
                .replaceAll("e", "📛")
                .replaceAll("f", "🎏")
                .replaceAll("g", "🐉")
                .replaceAll("ğ", "🐌")
                .replaceAll("h", "⛄")
                .replaceAll("ı", "📏")
                .replaceAll("i", "✏")
                .replaceAll("j", "🎷")
                .replaceAll("k", "🎋")
                .replaceAll("l", "🕒")
                .replaceAll("m", "👓")
                .replaceAll("n", "👠")
                .replaceAll("o", "📯")
                .replaceAll("ö", "🌞")
                .replaceAll("p", "🎧")
                .replaceAll("r", "💃")
                .replaceAll("s", "🐍")
                .replaceAll("ş", "🐢")
                .replaceAll("t", "☔")
                .replaceAll("u", "🔧")
                .replaceAll("ü", "🍇")
                .replaceAll("v", "☑")
                .replaceAll("y", "🎌")
                .replaceAll("z", "⚡")
                .replaceAll(" ", "     ");

        return kucuklatin;
    }

    public static String latinConverter(String andeyo) {
        andeyo = andeyo
                .replaceAll("🎄", "a")
                .replaceAll("🔩", "b")
                .replaceAll("🌜", "c")
                .replaceAll("💭", "ç")
                .replaceAll("🌛", "d")
                .replaceAll("📛", "e")
                .replaceAll("🎏", "f")
                .replaceAll("🐉", "g")
                .replaceAll("🐌", "ğ")
                .replaceAll("⛄", "h")
                .replaceAll("📏", "ı")
                .replaceAll("✏", "i")
                .replaceAll("🎷", "j")
                .replaceAll("🎋", "k")
                .replaceAll("🕒", "l")
                .replaceAll("👓", "m")
                .replaceAll("👠", "n")
                .replaceAll("📯", "o")
                .replaceAll("🌞", "ö")
                .replaceAll("🎧", "p")
                .replaceAll("💃", "r")
                .replaceAll("🐍", "s")
                .replaceAll("🐢", "ş")
                .replaceAll("☔", "t")
                .replaceAll("🔧", "u")
                .replaceAll("🍇", "ü")
                .replaceAll("☑", "v")
                .replaceAll("🎌", "y")
                .replaceAll("⚡", "z")
                .replaceAll("     ", " ");

        return andeyo;
    }

    public static boolean latinmi(String[] arr, String targetValue) {

        targetValue
                = targetValue.toLowerCase();
        targetValue
                = String.valueOf(targetValue.charAt(0));
        System.out.println(targetValue);
        return Arrays.asList(arr).contains(targetValue);
    }

}
