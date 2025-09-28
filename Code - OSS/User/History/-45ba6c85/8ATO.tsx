export default function ModalAccount(){
    return(
        <div className="absolute bg-dark-gray w-65 h-55">
            <div className="font-noto font-bold text-light-gray text-sm">

                <div className="m-3">
                    <p>CONTA</p>

                    <div className="flex mt-2">
                        <div className="h-15 w-15 rounded-full bg-light-gray"></div>
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