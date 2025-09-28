import Image from "next/image"

import person from "../../../public/logos/person.svg"

export default function ModalAccount(){
    return(
        <div className="absolute bg-dark-gray w-70 h-52">
            <div className="font-noto font-bold text-light-gray text-sm">

                <div className="m-3">
                    <p>CONTA</p>

                    <div className="flex mt-2">
                        <div className="h-14 w-14 content-center align-middle rounded-full bg-light-gray">
                            <Image src={person} alt=""></Image>
                        </div>
                        <div className="ml-3 text-xs content-center">
                            <p>Nome Conta</p>
                            <p className="font-normal">email@email.com</p>
                        </div>
                    </div>
                </div>
               
                <hr className="w-full h-0.5 rounded-sm"/>

                <button className="m-3 cursor-pointer">
                    <p>ALTERAR CONTA</p>
                </button>

                <hr className="w-full h-0.5 rounded-sm"/>

                <button className="m-3 cursor-pointer">
                    <p>LOGOUT</p>
                </button>
            </div>

        </div>
    )
}