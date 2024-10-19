import {defineStore} from 'pinia';
import {User} from "@/models/user";
import {v4 as uuidv4} from 'uuid';

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null as User | null,
    }),
    actions: {
        loadUser() {
            const savedUser = localStorage.getItem('user');
            if (savedUser) {
                this.user = JSON.parse(savedUser);
                console.log('load user from local storage', savedUser);
                return;
            }
            console.log('load user from local storage: no user found');
        },
        createUser: function ( name: string) {
            const id = uuidv4()
            this.user = new User(id, name);
            localStorage.setItem('user', JSON.stringify(this.user)); // TODO move to separate function
        },
    },
    getters: {
        userName(): string {
            return this.user?.name || '';
        }
    }
});