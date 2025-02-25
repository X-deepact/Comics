export interface NavigationChild {
    name: string;
    href: string;
    active?: boolean;
}

export interface NavigationItem {
    title: string;
    icon: string;
    children: NavigationChild[];
}
